package utils

import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "gorm.io/gorm/schema"
  "log"
  "os"
  "time"
)

func Open() (*gorm.DB, error) {
  // io writer
  writer := log.New(os.Stdout, "\r\n", log.LstdFlags)
  loggerConfig := logger.Config{
    SlowThreshold:             time.Second, // Slow SQL threshold
    LogLevel:                  logger.Info, // Log level
    IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
    Colorful:                  false,       // Disable color
  }

  newLogger := logger.New(writer, loggerConfig)

  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := "root:robot_123456#@tcp(192.168.3.9:3306)/mybatis_plus_study?charset=utf8mb4&parseTime=True&loc=Local"
  //db *gorm.DB
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    Logger: newLogger,
    NamingStrategy: schema.NamingStrategy{
      TablePrefix:   "api_",
      SingularTable: true,
    },
  })
  return db, err
}
