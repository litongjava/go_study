package main

import (
  "go-db-gorm-demo01/model"
  "gorm.io/driver/sqlite"
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
  // Globally mode
  db, err := gorm.Open(sqlite.Open("D:\\sqllite\\test01.db"), &gorm.Config{
    Logger: newLogger,
  })
  if err != nil {
    panic(err)
  }
  db.AutoMigrate(&model.Product{})
  print("success")
}
