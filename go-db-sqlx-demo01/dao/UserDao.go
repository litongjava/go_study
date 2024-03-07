package dao

import (
  "fmt"
  "go-db-sqlx-demo01/container"
  "go-db-sqlx-demo01/models"
)

// 更新数据
func updateRowDemo() {
  sqlStr := "update user set age=? where id = ?"
  ret, err := container.Db.Exec(sqlStr, 39, 6)
  if err != nil {
    fmt.Printf("update failed, err:%v\n", err)
    return
  }
  n, err := ret.RowsAffected()
  if err != nil {
    fmt.Printf("get RowsAffected failed, err:%v\n", err)
    return
  }
  fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo() {
  sqlStr := "delete from student where id = ?"
  ret, err := container.Db.Exec(sqlStr, 6)
  if err != nil {
    fmt.Printf("delete failed, err:%v\n", err)
    return
  }
  n, err := ret.RowsAffected()
  if err != nil {
    fmt.Printf("get RowsAffected failed, err:%v\n", err)
    return
  }
  fmt.Printf("delete success, affected rows:%d\n", n)
}

func selectNamedQuery() {
  sqlStr := "select * from student where age = :age"
  args := map[string]interface{}{
    "age": 22,
  }
  rows, err := container.Db.NamedQuery(sqlStr, args)
  if err != nil {
    fmt.Printf("named query failed failed, err:%v\n", err)
    return
  }
  defer rows.Close()
  for rows.Next() {
    var u models.User
    if err := rows.StructScan(&u); err != nil {
      fmt.Printf("struct sacn failed, err:%v\n", err)
      continue
    }
    fmt.Printf("%#v\n", u)
  }
}

func insertUserDemo() {
  users := []models.User{
    {Name: "111", Age: 18},
    {Name: "222", Age: 18},
    {Name: "333", Age: 18},
  }
  sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
  _, err := container.Db.NamedExec(sqlStr, users)
  if err != nil {
    panic(err)
  }
}
