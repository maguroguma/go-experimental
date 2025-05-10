## 前提: エラーの wrap

以下のように、`fmt.Errorf` によって、エラーをラップしたエラーというのが作れる。

```go
	err := fmt.Errorf("wrap: %w", os.ErrNotExist)
```

あるいは、`Unwrap()` メソッドを実装したカスタム型でも、ラップしたエラーというのは実現できる。

## `errors.Is`: あるエラーが含まれる「かどうか」を検査する

以下のように、エラーごとに処理を分けることができる。
引数には、「検査対象のエラー」と、「含まれるか検証したいエラーの **実体** 」を渡す。

```go
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("ファイルが存在しません")
	}
	if errors.Is(err, os.ErrPermission) {
		fmt.Println("ファイルのアクセス権限がありません")
	}
```


true or false のみが知りたいだけなら、`errors.Is` を使えばよい。

## `errors.As`: エラーを特定の型に変換する

こちらもエラーがラップされているか、というのを調べるものになるが、雰囲気としては型アサーションに近い。
引数には、「検査対象のエラー」と、「含まれるか検証したいエラーのポインタ型[^1]」を渡す。

```go
	err := fmt.Errorf("wrap: %w", &MyError{Code: 404})

	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Printf("特定のエラーが発生: %d\n", myErr.Code)
	}
```

[^1]: 雰囲気としては、型そのものを渡しているに近い。

## まとめ

アプリケーション開発者的には、Is のほうが圧倒的に使う機会が多そう（な気がする）。
