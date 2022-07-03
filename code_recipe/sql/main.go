package main

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("pgx", "host=localhost port=5432 user=testuser dbname=testdb password=pass sslmode=disable")
	if err != nil {
		// エラーハンドリング
	}

	err = db.PingContext(ctx)
	if err != nil {
		// エラーハンドリング
	}
}
