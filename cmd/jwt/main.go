package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// シークレットキー（実際の運用では環境変数などで安全に管理すべき）
var jwtSecretKey = []byte("my_super_secret_key_for_jwt_signing")

// カスタムクレーム - JWT のペイロード部分
type CustomClaims struct {
	Username             string `json:"username"`
	Role                 string `json:"role"`
	jwt.RegisteredClaims        // これを埋め込むことで claim interface を実装したことになる
}

// JWT トークンを生成する関数
func generateJWT(username, role string, expirationMinutes int) (string, error) {
	// トークンの有効期限を設定
	expirationTime := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)

	// カスタムクレームを作成
	claims := &CustomClaims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			// トークンが有効になる時刻
			NotBefore: jwt.NewNumericDate(time.Now()),
			// トークンの発行時刻
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// トークンの有効期限
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			// トークン発行者
			Issuer: "jwt-demo-service",
			// トークンの対象者（ユーザーを特定）
			Subject: username,
			// トークンの一意の識別子（オプション）
			ID: "1234567890",
		},
	}

	// 新しいトークンを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// JWT トークンを検証する関数
func validateJWT(tokenString string) (*CustomClaims, error) {
	// トークンを解析して検証
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 署名アルゴリズムの検証
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("予期しない署名方式: %v", token.Header["alg"])
		}
		// 検証に使用するシークレットキーを返す
		return jwtSecretKey, nil
	})

	// エラーがあれば返す
	if err != nil {
		return nil, err
	}

	// トークンが有効かどうか確認
	if !token.Valid {
		return nil, fmt.Errorf("無効なトークン")
	}

	// トークンからクレームを取得
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("クレームの型変換に失敗")
	}

	return claims, nil
}

func main() {
	// ===== JWT トークンの生成 =====
	username := "user123"
	role := "admin"
	tokenExpirationMinutes := 60 // 60分後に有効期限が切れる

	// トークンを生成
	tokenString, err := generateJWT(username, role, tokenExpirationMinutes)
	if err != nil {
		log.Fatalf("トークン生成エラー: %v", err)
	}
	fmt.Println("生成されたJWTトークン:")
	fmt.Println(tokenString)
	fmt.Println()

	// ===== JWT トークンの検証 =====
	// 有効なトークンを検証
	fmt.Println("トークン検証結果:")
	claims, err := validateJWT(tokenString)
	if err != nil {
		fmt.Printf("トークン検証エラー: %v\n", err)
	} else {
		fmt.Println("トークン検証成功!")
		fmt.Printf("ユーザー名: %s\n", claims.Username)
		fmt.Printf("ロール: %s\n", claims.Role)
		fmt.Printf("有効期限: %v\n", claims.ExpiresAt)
	}
	fmt.Println()

	// ===== 無効なトークンの例 =====
	invalidToken := tokenString + "invalid"
	fmt.Println("無効なトークン検証結果:")
	_, err = validateJWT(invalidToken)
	if err != nil {
		fmt.Printf("予期された検証エラー: %v\n", err)
	} else {
		fmt.Println("このトークンは無効なはずなのに検証に成功しました")
	}
	fmt.Println()
}
