package main

import (
  "go-db-gorm-demo01/model"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "log"
  "os"
  "time"
)

func main() {
  newLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
    logger.Config{
      SlowThreshold:             time.Second,   // Slow SQL threshold
      LogLevel:                  logger.Silent, // Log level
      IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
      Colorful:                  false,         // Disable color
    },
  )

  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := "root:robot_123456#@tcp(192.168.3.9:3306)/mybatis_plus_study?charset=utf8mb4&parseTime=True&loc=Local"
  //db *gorm.DB
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    Logger: newLogger,
  })
  if err != nil {
    panic(err)
  }
  print(db)

  err = db.AutoMigrate(&model.Product{})
  if err != nil {
    panic(err)
  }

}
