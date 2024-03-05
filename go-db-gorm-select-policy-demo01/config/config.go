package config

import (
  _ "github.com/joho/godotenv/autoload"
  "os"
)

var MODE = os.Getenv("MODE")
var DATABASE_DSN = os.Getenv("DATABASE_DSN")
var DATABASE_REPLICAS = os.Getenv("DATABASE_REPLICAS")
