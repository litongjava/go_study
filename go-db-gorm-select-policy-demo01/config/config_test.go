package config

import (
  "fmt"
  _ "github.com/joho/godotenv/autoload"
  "os"
  "testing"
)

func TestGetKey(t *testing.T) {
  // 尝试显式加载.env文件
  fmt.Println("start")
  fmt.Println(os.Getenv("DATABASE_DSN"))
  fmt.Println("end")
}

func TestGetDATABASEDSN(t *testing.T) {
  print(DATABASE_DSN)
}
