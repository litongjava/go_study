package main

import (
  "fmt"
  "go-db-sqlx-demo01/container"
  "go-db-sqlx-demo01/sqlite"
  "testing"
)

func TestInsert(t *testing.T) {
  sqlite.OpenDb()
  // 插入数据
  sqlStr := "insert into user (name, age) values (?,?)"
  ret, err := container.Db.Exec(sqlStr, "Da Li", 18)
  if err != nil {
    fmt.Printf("insert failed, err:%v\n", err)
    return
  }
  id, err := ret.LastInsertId()
  if err != nil {
    fmt.Printf("get lastinsert ID failed, err:%v\n", err)
    return
  }
  fmt.Printf("insert success, the id is %d.\n", id)
}
