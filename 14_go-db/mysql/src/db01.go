package main

import (
  "context"
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "log"
)

var db *sql.DB

const (
  server   = "127.0.0.1"
  port     = 3306
  user     = "root"
  password = "123456"
  database = "go_db"
)

func init() {
  //设置Flats为 日期 时间 微秒 文件名:行号
  log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}
func test01() {
  //format := "server=%s;port=%d;user id=%s;password=%s;database=%s;"
  format := "%s:%s@tcp(%s:%d)/%s"
  var connStr = fmt.Sprintf(format, user, password, server, port, database)
  fmt.Println(connStr)
  var err error
  db, err = sql.Open("mysql", connStr)
  if err != nil {
    log.Fatalln(err)
  }

  var ctx context.Context = context.Background()
  err = db.PingContext(ctx)
  if err != nil {
    log.Fatalln(err)
  }
  fmt.Println("connected")

  //many, err := getMany(1)
  //if err != nil {
  //  log.Fatalln(err)
  //}
  //fmt.Println(many)

  //one, err := getOne(1)
  //if err != nil {
  //  log.Fatalln(err)
  //}
  //fmt.Println(one)

  //one.name += "001"
  //one.Update()
  //if err != nil {
  //  log.Fatalln(err.Error())
  //}

  //one, err = getOne(1)
  //if err != nil {
  //  log.Fatalln(err)
  //}
  //fmt.Println(one)

  //err = one.Delete()
  //if err != nil {
  //  log.Fatalln(err)
  //}
  a := app{
    name:   "Test",
    order:  1123,
    level:  10,
    status: 1,
  }
  err = a.Insert()
  if err != nil {
    log.Fatalln(err)
  }
}
