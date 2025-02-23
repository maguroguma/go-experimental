最悪、非同期処理を書くのが下手でも context パッケージは知ってないと Go が書けない、と思う。

[zenn book - よくわかるcontextの使い方](https://zenn.dev/hsaki/books/golang-context)を復習する。

## 事前知識

### chan の close 時の挙動について

- close chan の受信は問題なくできるし、正しく受けとれたかどうかの確認もできる
    - close chan がバッファにデータを持っているときは、バッファ分を通常受信してあと、ゼロ値の受信となる
- close chan への送信はできない、panic になる
    - 毎回忘れてしまうが、受信側はゼロ値の受信で自然にできるので、そこから覚えたい
    - 送信する側が「その chan が close しているかどうか？」を知る手段は、恐らく標準的なものは用意されていない
- chan を close することによって、受信待ちを解除できるため、close でもブロックの解除ができる、と考えてよい（はず）
- for range chan の構文により、close まで自然にループさせることができる
- select case の構文でも、close chan は検知できる
- **closed chan は複数の goroutine で検知できる**
    - context を理解する上で重要！
- close chan を close すると panic になる！
    - context の cancel 関数はこの心配がないので便利！

GPT と会話していたら「送信者は close しないのが基本、受信者が close する」と言われたが、
これはまだ疑問なので保留にしておく。

### chan のデフォルトサイズ `make(chan int)` について

> ## 🙋🙋🙋 User 🙋🙋🙋───
> 
> go の make(chan int) で作った chan のサイズって1で合ってますか？
> 
> ## 🤖🤖🤖 Copilot 🤖🤖🤖───
> 
> いいえ、`make(chan int)` で作成したチャネルのサイズは 0 です。バッファなしのチャネルになります。サイズ 1 のチャネルを作成するには、`make(chan int, 1)` と指定します。
> 
> ## 🙋🙋🙋 User 🙋🙋🙋───
> 
> バッファ 0 の chan ってデータの送受信は可能なのでしょうか？
> 
> ## 🤖🤖🤖 Copilot 🤖🤖🤖───
> 
> はい、バッファ 0 のチャネルでもデータの送受信は可能です。ただし、送信側と受信側が同時に準備できている必要があります。送信は受信が準備されるまでブロックされ、受信は送信が準備されるまでブロックされます。これにより、同期的な通信が実現されます。

実験したが、これは合っているっぽい。
chan の最小サイズは1だと勘違いしてしまっていた。

サイズ 0 の挙動は、完全に送信と受信が同期しており、chan はただのデータの通り道のようになる、と理解したい
（chan 内部でデータのバッファリングを一切行わない）。

## 意識改革

- Web サーバプロセスや DB リクエストなど、時間がかかる処理は内部で goroutine が作成されている
    - DB リクエストに関しては、1つのリクエストで複数の DB レプリカに複数の goroutine で問い合わせる、といった設計もありうる、らしい
    - 親関数から木構造で goroutine が分岐するイメージを持つと良いのかもしれない
- 情報伝達の基本は引数に chan を渡すことだけど、それだと色々と設計に難が出てきてしまう
    - **関数の第一引数に `context.Context` を1つ渡すだけで、色々とおいしい思いが出来る**、というのが context の基本概念っぽい
- `<-chan struct{}` という受信専用の chan 型は、close によってのみ通知される用途を示唆していると言ってよい（のかもしれない）
- context を受け入れる関数は、done チャネルから受信できることで、キャンセル処理を実装できる（多分）
    - done チャネルは cancel 関数を呼ばれると close されるため、それを検知して後処理を行う
    - **「cancel 関数の実行」と「chan の close」を等価なものとみなしたい**
- 同一 context から生成した兄弟 context については、それぞれ別の Done, cancel に別れる
- **一方で、親子関係の context は、親の cancel を呼ぶと子孫の context の Done チャネルが閉じられる**
    - 逆方向のキャンセルはできない
- `context.WithDeadline` 関数を使うと、「指定した時刻で」自動的に Done チャネルが閉じられる
    - 一緒に生成される cancel 関数で明示的に閉じることもできる
- `context.WithTimeout` は時間バージョン
- **タイムアウトされたあとに cancel をコールしても問題ない、むしろするべきとされている**
    - close chan を close すると panic だが、cancel は何度呼んでも問題ない
- Err メソッドにより、キャンセルされていない、明示的なキャンセルがされた、タイムアウトした、が判断できる
    - Go 1.21 から追加された Cause という概念により、cancel 時にエラーをカスタムして報告し、Done で通知を受けた側がそのエラーを取り出すことが出来るようになった
- `context.WithValue` 関数によって、context に値を持たせることが出来る
    - 取り出すときは `interface{}` 型になる、引数を context に好き放題隠蔽できる、といった見方もできて難しいので、実用的な例で別に学びたい

## Context

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

## context を受け取る関数について解釈する

context を受け取るライブラリ側の関数は、恐らくそのなかで Done の検知を行っている。
もしも親側、すなわち自分達が書いたコード側で cancel を実行したとき、ライブラリ側では適切なキャンセル処理を行ってくれている、はず（実験してみたい）。

自動的に、context を受け取る関数は、時間がかかる処理を行っており、何らかの goroutine を起動していることを示唆している、はず。

「context を必要とするなら、第一引数に明示するべし」というのは↓でも言及されている。
基本的に、構造体のフィールドに context を持たせるのはやめたほうが良さそう。  
https://zenn.dev/hsaki/books/golang-context/viewer/pkgdesign#ok%E4%BE%8B%3A-myfunc%E3%81%AE%E7%AC%AC%E4%B8%80%E5%BC%95%E6%95%B0%E3%81%ABcontext%E3%82%92%E8%BF%BD%E5%8A%A0

## 次にやってみたいこと

[tenntenn さんのハンズオン](https://github.com/gohandson/goroutine-ja)があるので、余裕があればやってみる。

