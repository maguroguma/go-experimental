# defer について

## defer の処理順序

あとから登録したものほど先に実行される。LIFO。

> このような設計になっているのは、defer が何かしらの「後処理」に使われることが多いため。
> 複数のリソースを開いた場合、後から開いたものを先に閉じるようにするのが自然ではある。

AI からはこういった回答だったけど…これは後付けのような気がする？
「defer が入れ子になっているとき、スタックを利用するほうが都合が良い」という回答もあり、こちらのほうが自然に感じられる。

入れ子になっていようがいまいが、スタックを用いて管理されるから自然と LIFO になる、と覚えておきたい。

`defer f()` で `f` を予約した、みたいなイメージだったが、スタックに積んだ、という風に考えたほうが色々と合点がいくことが多そう。

## 到達しなかった defer

スタックに積まれない、実行されない、と考えてよい。

## defer による return の上書き

以下のよう名前付き返り値を使うと defer で結果を上書きできる。
エラー処理とかに使える場面はありそう。

```go
func overwrite() (result int) {
	defer func() {
		result = 42 // return の値を書き換える
	}()
	return 1
}
```

## panic 時も登録済みのものは実行される

panic 時にも、panic するまでに登録されたものは、ちゃんと実行される。
なので、以下のような defer を登録しておくことで、panic から recover することができる。

```go
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered in testWhenPanic: %v\n", r)
		}
	}()
```

## 他言語での defer

具体的には JavaScript で defer が欲しくなるときがある。
この場合は `try finally` で実現できるっぽい。

また、以下のような関数を自作すれば defer っぽいことはできるようになる。

```js
// fn に defer という関数を渡しつつ、代理で fn を実行し結果を得る
function withDefer(fn) {
  const deferred = []; // defer スタック
  const defer = (callback) => deferred.push(callback);

  const result = fn(defer); // 本来実行したい処理を実行する

  while (deferred.length) {
    deferred.pop()(); // LIFO で実行
  }

  return result;
}

// 例: 複数の return がある関数
function deferExample() {
  return withDefer((defer) => {
    console.log("Start");

    defer(() => console.log("Cleanup 1st"));

    if (Math.random() > 0.5) {
      console.log("Early return 1");
      return "Result 1";
    }

    defer(() => console.log("Cleanup 2nd"));

    console.log("Early return 2");
    return "Result 2";
  });
}
```

思いがけずに Promise のパターンを理解する良い題材になってる…かもしれない。

高階関数は「関数を引数にとる」「関数を返す」「クロージャ」あたりを頭に入れておくのが理解のポイントな気がする。

あと、関数を渡した時点で「あとから代理で実行される」と思うべき。
渡す関数が謎の関数引数を持っているとしたら、その代理で実行する主体が「関数実体である何か」を渡しつつ実行する、と思うのが良さそう。
「あとから」の部分は、遅延実行される場合もあるので、曖昧に捉えるのが良いと思う。

### これはダブルディスパッチなのか？

[ダブル・ディスパッチ～ 典型的なオブジェクト指向プログラミング・イディオム ～](https://www.infoq.com/jp/articles/DoubleDispatch_0829/)

オブジェクト指向な言語におけるダブルディスパッチの概略↓。

```
public void someMethod( SomeObject parameter ) {
    parameter.otherMethod( this ) ;
}
```

登場するオブジェクト2つを関数だと考えると、そのような類推も出来なくもない気もするが、無理に当てはめるとわけが分からなくなりそう。

メッセージ通信の流れを捉えると分かりやすいのかもしれない。

> someMethod（を持つオブジェクト A）へメッセージ送信 → A が parameter オブジェクト（B）へ自身を渡しつつメッセージ送信 → B が応答する
