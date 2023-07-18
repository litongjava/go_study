package main

import (
  "database/sql"
  "fmt"
  "log"
)

func getOne(id int) (a app, err error) {
  var row *sql.Row = db.QueryRow("select id,name,status,level,`order` from app where id =?", id)
  err = row.Scan(&a.id, &a.name, &a.status, &a.level, &a.order)
  return
}

func getMany(status int) (apps []app, err error) {
  var sql = "select id,name,status,level,`order` from app where status=?"
  rows, err := db.Query(sql, status)
  for rows.Next() {
    a := app{}
    err = rows.Scan(&a.id, &a.name, &a.status, &a.level, &a.order)
    if err != nil {
      log.Fatalln(err)
    }
    apps = append(apps, a)
  }
  return
}
func (a *app) Update() (err error) {
  sql := "UPDATE app set name=? where id=?"
  exec, err := db.Exec(sql, a.name, a.id)
  fmt.Println(exec)
  return
}

func (a *app) Delete() (err error) {
  sql := "DELETE from app where id=?"
  exec, err := db.Exec(sql, a.id)
  fmt.Println(exec)
  return
}

func (a *app) Insert() (err error) {
  var insertSql = "insert into app(name,status,level,`order`) values (?,?,?,?);"
  //var maxId = "select max(id) from app;"
  //var statementSql = insertSql + maxId
  //var statementSql = "insert into app(name,status,level,`order`) values ('Test',1,1,1);select max(id) from app;"
  statement, err := db.Prepare(insertSql)
  if err != nil {
    log.Fatalln(err)
  }
  defer statement.Close()
  statement.QueryRow(a.name, a.status, a.level, a.order)
  return
}
