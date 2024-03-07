package main

import (
  "fmt"
  "go-db-sqlx-demo01/container"
  "go-db-sqlx-demo01/models"
  "go-db-sqlx-demo01/sqlite"
  "testing"
)

func TestSinleQuery(t *testing.T) {
  sqlite.OpenDb()
  // 单条查询
  sqlStr := "select * from user where id = ?"
  var u models.User
  if err := container.Db.Get(&u, sqlStr, 1); err != nil {
    return
  }
  fmt.Printf("id:%d,name:%s,age:%d", u.Id, u.Name, u.Age)
}

// 多条查询
func queryMultiRow() []models.User {
  sqlStr := "select * from user"
  var users []models.User
  if err := container.Db.Select(&users, sqlStr); err != nil {
    return nil
  }
  fmt.Println(users)
  return users
}
