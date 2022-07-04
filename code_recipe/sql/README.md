# sql

## データベースの基本的な利用法

アプリケーションコード（ORMを含む）
↓
汎用API（database/sqlパッケージ）
↓
データベースドライバー
↓
データベース

コネクションプーリングや並行処理の制御のような処理は`database/sql`パッケージで吸収する。  
DBアクセスのときは対象のドライバーをimportして利用する。  
ORMは標準ライブラリにはない。  
サードパーティには`sqlx`,`Gorm`,`ent`など。それらも`database/sql`パッケージのラッパー。

## データベースへ接続

```sh
docker run -d --name go-sample-postgres -e POSTGRES_USER=testuser -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=testdb -p 5432:5432 postgres
```

ドライバーのダウンロード（PostgreSQL, jackc/pgx）

```sh
go get github.com/jackc/pgx/v4
# バージョン指定
# go get github.com/jackc/pgx/v4@v4.1.0
```

## ドライバーのブランクインポート

ドライバーを用いる場合、ブランクインポートしてドライバーの初期化のみを実施する方法がよくある

```go
import (
  "database/sql"
  _ "github.com/jackc/pgx/v4/stdlib" // ブランクインポートで初期化のみ実施
)
```

同じデータベースにアクセスする場合、sql.Open関数を用いて生成したsql.DB構造体を使い回すことになる。  
init()関数内や、main関数に近い場所で呼ばれることが多い。

データベースへコネクションが確立しているかどうかを確かめる。

```go
	err = db.Ping()
	if err != nil {
		// エラーハンドリング
	}
```

`database/sql`パッケージに含まれる関数やメソッドの多くはコンテキストを扱える。（XxxContextという名前のもの）  
タイムアウトやキャンセルを設定できるため、XxxContextの関数やメソッドを使っていくのが良い。

## pgxを使用したクエリーのロギング

