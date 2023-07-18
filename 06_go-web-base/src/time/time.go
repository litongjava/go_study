package main

import (
  "fmt"
  "time"
)

func main() {
  var now time.Time = time.Now()
  fmt.Println(now)
  var unix int64 = now.Unix()
  fmt.Println(unix)
}
