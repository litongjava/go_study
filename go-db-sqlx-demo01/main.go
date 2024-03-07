package main

import (
  "fmt"
  "go-db-sqlx-demo01/sqlite"
)

func main() {
  err := sqlite.OpenDb()
  if err != nil {
    panic(err)
  }
  fmt.Println("Hello,World")
}
