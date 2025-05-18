package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// シークレットキー（実際の運用では環境変数などで安全に管理すべき）
var jwtSecretKey = []byte("my_super_secret_key_for_jwt_signing")

// ユーザー認証情報（実際の運用ではデータベースなどを使用）
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
	"admin": "admin123",
}

// ユーザーロール（実際の運用ではデータベースなどを使用）
var userRoles = map[string]string{
	"user1": "user",
	"user2": "user",
	"admin": "admin",
}

// カスタムクレーム - JWT のペイロード部分
type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// ログインリクエスト
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ログインレスポンス
type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresIn int    `json:"expiresIn"` // 有効期間（秒）
	Username  string `json:"username"`
	Role      string `json:"role"`
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
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "jwt-demo-api",
			Subject:   username,
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
		return jwtSecretKey, nil
	})

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

// ログインハンドラー - ユーザー名とパスワードを検証してJWTトークンを発行
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// POST以外のメソッドを拒否
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// リクエストからJSONを解析
	var loginReq LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// ユーザー認証
	storedPassword, exists := users[loginReq.Username]
	if !exists || storedPassword != loginReq.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// ユーザーロールを取得
	role := userRoles[loginReq.Username]

	// JWTトークンを生成（有効期限は60分）
	tokenString, err := generateJWT(loginReq.Username, role, 60)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// レスポンスを作成
	response := LoginResponse{
		Token:     tokenString,
		ExpiresIn: 60 * 60, // 秒単位（60分）
		Username:  loginReq.Username,
		Role:      role,
	}

	// JSONレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// 保護されたリソースへのアクセスを制御するミドルウェア
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Authorizationヘッダーを取得
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Bearer トークンの形式を確認
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// トークン部分を抽出
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// トークンを検証
		claims, err := validateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// ユーザー情報をリクエストコンテキストに追加（簡易的な実装）
		r.Header.Set("X-User", claims.Username)
		r.Header.Set("X-Role", claims.Role)

		// 次のハンドラーを実行
		next(w, r)
	}
}

// 保護されたリソースにアクセスするハンドラー
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-User")
	role := r.Header.Get("X-Role")

	// レスポンスを返す
	response := map[string]string{
		"message":  "保護されたリソースにアクセスしました",
		"username": username,
		"role":     role,
		"time":     time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 管理者専用のハンドラー
func adminOnlyHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("X-Role")
	if role != "admin" {
		http.Error(w, "管理者権限が必要です", http.StatusForbidden)
		return
	}

	response := map[string]string{
		"message": "管理者向け保護リソースにアクセスしました",
		"time":    time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// インデックスページのためのハンドラー
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>JWT デモ</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
        .container { margin-bottom: 20px; }
        button { padding: 8px 16px; background-color: #4CAF50; color: white; border: none; cursor: pointer; margin-right: 5px; }
        input { padding: 8px; margin-bottom: 10px; width: 200px; }
        pre { background-color: #f5f5f5; padding: 10px; overflow-x: auto; }
        .output { border: 1px solid #ddd; padding: 10px; min-height: 50px; margin-top: 10px; }
    </style>
</head>
<body>
    <h1>JWT 認証デモ</h1>
    
    <div class="container">
        <h2>1. ログイン (JWT トークン取得)</h2>
        <div>
            <input type="text" id="username" placeholder="Username" value="user1"><br>
            <input type="password" id="password" placeholder="Password" value="password1"><br>
            <button onclick="login()">ログイン</button>
        </div>
        <pre class="output" id="loginOutput">// ここにログイン結果が表示されます</pre>
    </div>

    <div class="container">
        <h2>2. 保護されたリソースへのアクセス</h2>
        <button onclick="accessProtected()">保護リソースにアクセス</button>
        <pre class="output" id="protectedOutput">// ここに保護リソースの結果が表示されます</pre>
    </div>

    <div class="container">
        <h2>3. 管理者専用リソースへのアクセス</h2>
        <button onclick="accessAdminOnly()">管理者リソースにアクセス</button>
        <pre class="output" id="adminOutput">// ここに管理者リソースの結果が表示されます</pre>
    </div>

    <div class="container">
        <h2>4. トークン情報</h2>
        <button onclick="decodeToken()">トークンをデコード</button>
        <pre class="output" id="tokenInfo">// ここにトークン情報が表示されます</pre>
    </div>

    <script>
        let token = '';
        
        async function login() {
            const username = document.getElementById('username').value;
            const password = document.getElementById('password').value;
            
            try {
                const response = await fetch('/api/login', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({ username, password })
                });
                
                const result = await response.json();
                if (response.ok) {
                    token = result.token;
                    document.getElementById('loginOutput').textContent = JSON.stringify(result, null, 2);
                } else {
                    document.getElementById('loginOutput').textContent = 'エラー: ' + result.error;
                }
            } catch (err) {
                document.getElementById('loginOutput').textContent = 'エラー: ' + err.message;
            }
        }
        
        async function accessProtected() {
            if (!token) {
                document.getElementById('protectedOutput').textContent = 'エラー: 先にログインしてください';
                return;
            }
            
            try {
                const response = await fetch('/api/protected', {
                    headers: {
                        'Authorization': 'Bearer ' + token
                    }
                });
                
                const result = await response.json();
                document.getElementById('protectedOutput').textContent = JSON.stringify(result, null, 2);
            } catch (err) {
                document.getElementById('protectedOutput').textContent = 'エラー: ' + err.message;
            }
        }
        
        async function accessAdminOnly() {
            if (!token) {
                document.getElementById('adminOutput').textContent = 'エラー: 先にログインしてください';
                return;
            }
            
            try {
                const response = await fetch('/api/admin', {
                    headers: {
                        'Authorization': 'Bearer ' + token
                    }
                });
                
                if (response.ok) {
                    const result = await response.json();
                    document.getElementById('adminOutput').textContent = JSON.stringify(result, null, 2);
                } else {
                    const result = await response.text();
                    document.getElementById('adminOutput').textContent = 'エラー: ' + result;
                }
            } catch (err) {
                document.getElementById('adminOutput').textContent = 'エラー: ' + err.message;
            }
        }
        
        function decodeToken() {
            if (!token) {
                document.getElementById('tokenInfo').textContent = 'エラー: 先にログインしてください';
                return;
            }
            
            try {
                const parts = token.split('.');
                if (parts.length !== 3) {
                    throw new Error('無効なトークン形式');
                }
                
                const header = JSON.parse(atob(parts[0]));
                const payload = JSON.parse(atob(parts[1]));
                
                const info = {
                    header: header,
                    payload: payload,
                    // 署名部分は検証用なのでデコードしません
                    signature: '(省略)...'
                };
                
                document.getElementById('tokenInfo').textContent = JSON.stringify(info, null, 2);
            } catch (err) {
                document.getElementById('tokenInfo').textContent = 'エラー: ' + err.message;
            }
        }
    </script>
</body>
</html>
`
	t, err := template.New("index").Parse(tmpl)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func main() {
	// ルーティング設定
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/protected", authMiddleware(protectedHandler))
	http.HandleFunc("/api/admin", authMiddleware(adminOnlyHandler))

	// サーバー起動
	fmt.Println("サーバーを起動します... http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
