package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var _ pgx.Logger = (*logger)(nil)

type logger struct{}

func (l *logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	if msg == "Query" {
		log.Printf("SQL: %v\nArgs: %v\n", data["sql"], data["args"])
	}
}

func main() {
	ctx := context.Background()

	config, err := pgx.ParseConfig("user=testuser password=pass host=localhost port=5432 dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatalf("parse config: %v\n", err)
	}
	config.Logger = &logger{}

	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatalf("connect: %v\n", err)
	}

	sql := `SELECT schemaname, tablename FROM pg_tables WHERE schemaname = $1;`
	args := `information_schema`

	rows, err := conn.Query(ctx, sql, args)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var pgtables []PgTable
	for rows.Next() {
		var s, t string
		if err != rows.Scan(&s, &t); err != nil {
			// エラー
		}
		pgtables = append(pgtables, PgTable{SchemaName: s, TableName: t})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
