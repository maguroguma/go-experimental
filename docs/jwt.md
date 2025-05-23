# JWT

## 基礎知識まとめ

`Header.Payload.Sign` という形式の文字列。

それぞれのパートは、もとのデータに対して Base64url エンコードを施したものになっている。
ヘッダーとペイロードは JSON であり、署名はヘッダーとペイロードを元に、シークレットキーを用いて作成される。

署名部分は、サーバ側で管理される秘密鍵によって行われる。
これにより、改竄を検知できる。

## 深掘り

### Base64url エンコード

Base64 の URL に使える文字種でエンコードされる版、と考えておく。

### 署名生成に利用するアルゴリズム

#### HS256 (HMAC-SHA256)

共通の秘密鍵を使用した対称暗号、すなわち、署名と検証を同じ鍵（シークレット）を用いて行う。
実装が簡単かつ高速なのが利点となっている。

##### HMAC って？

Hash-based Message Authentication Code の略。
HMACは「このメッセージは、秘密鍵を持っている人が作成し、それ以降変更されていない」ことを証明する仕組み。

「ハッシュ関数」と「秘密鍵」を組み合わせて成立するもの。
ハッシュ関数としては SHA256 がよく使われる、すなわち、HS256 が良く使われる。

メリットとして「高速処理が可能、比較的シンプルな実装、暗号学的に安全」といった点がある。

式にすると以下のようになる。

```txt
HMAC(K, m) = H((K ⊕ opad) || H((K ⊕ ipad) || m))
```

H がハッシュ関数で K が秘密鍵 で m がメッセージ。
ハッシュ関数は一方向なので、たとえ秘密鍵を知っていても、HMAC の結果から元のメッセージを復元することはできない。
あくまで、目的は「暗号化」ではなく「改竄検知、認証」であることに注意する。

##### ハッシュ関数について

任意のバイト列を、アルゴリズムごとの固定長バイト列に変換する。
Go 的なシグニチャでいうと以下のようになる。
なんとなく考えていると、入力を文字列に限定して考えてしまいそうになるが、任意のバイナリデータをハッシュ化できる。

```go
func sha256(input []byte) []byte
```

当然だが、ハッシュ関数自体には秘密鍵がどうとかは無関係である。
以下のように、CLI でも実行できることを知っていると、そこらへんの混乱はなくなると思う。

```sh
# 文字列のハッシュ計算
echo -n "ハッシュ化したい文字列" | shasum -a 256

# またはファイルのハッシュ計算
shasum -a 256 ファイル名.txt
```

#### RS256 (RSA-SHA256)

公開鍵暗号を使用した非対称暗号、すなわち、署名と検証で異なる鍵（公開鍵と秘密鍵）を用いる。
署名には秘密鍵を使用し、検証には公開鍵を使用する。

マイクロサービスなど、複数のサービス間で共有する場合に適している。

### クレーム

「JWT のペイロードはクレームである」という説明になると思う。

RFC7519 で標準クレームが定義されているほか、プライベートクレームといってアプリケーション固有の情報を含められる。

