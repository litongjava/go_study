package main

import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func main() {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := "root:robot_123456#@tcp(192.168.3.9:3306)/enote?charset=utf8mb4&parseTime=True&loc=Local"
  //db *gorm.DB
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic(err)
  }
  print(db)
}
