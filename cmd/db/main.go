package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/maguroguma/go-experimental/sqlc/db"
)

func main() {
	connStr := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	ctx := context.Background()

	// ユーザー作成
	user, err := queries.CreateUser(ctx, "Alice")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created User:", user)

	// ユーザー取得
	fetchedUser, err := queries.GetUser(ctx, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetched User:", fetchedUser)
}
