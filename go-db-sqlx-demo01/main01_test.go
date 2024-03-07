package main

//import (
//  "fmt"
//  "github.com/jmoiron/sqlx"
//  _ "gorm.io/driver/mysql"
//  "testing"
//)
//
//var db *sqlx.DB
//
//func initDB() (err error) {
//  dsn := "root:robot_123456#@tcp(192.168.3.9:3306)/mybatis_plus_study?charset=utf8mb4&parseTime=True"
//  // 也可以使用MustConnect连接不成功就panic
//  db, err = sqlx.Connect("mysql", dsn)
//  if err != nil {
//    fmt.Printf("connect DB failed, err:%v\n", err)
//    return
//  }
//  db.SetMaxOpenConns(20)
//  db.SetMaxIdleConns(10)
//  return
//}
//func TestConnecSqlLite(test *testing.T) {
//  err := initDB()
//  if err != nil {
//    panic(err)
//  }
//  fmt.Println("Hello,World")
//}
