package main

import (
  "fmt"
  "go-db-sqlx-demo01/sqlite"
  "testing"
)

func TestConnectSqlite(test *testing.T) {
  err := sqlite.OpenDb()
  if err != nil {
    panic(err)
  }
  fmt.Println("Hello,World")
}
