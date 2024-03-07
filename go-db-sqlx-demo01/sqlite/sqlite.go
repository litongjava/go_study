package sqlite

import (
  "fmt"
  "github.com/jmoiron/sqlx"
  _ "github.com/mattn/go-sqlite3"
  "go-db-sqlx-demo01/container"
)

func OpenDb() (err error) {
  dsn := "file:E.db?cache=shared&mode=rwc" // Example SQLite DSN
  // Connect to the database using sqlx with the SQLite driver
  db, err := sqlx.Connect("sqlite3", dsn)
  if err != nil {
    fmt.Printf("connect DB failed, err:%v\n", err)
    return
  }
  // For SQLite, these settings are less relevant but setting them doesn't harm
  db.SetMaxOpenConns(20)
  db.SetMaxIdleConns(10)

  container.Db = db
  return
}
