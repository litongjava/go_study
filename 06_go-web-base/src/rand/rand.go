package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().Unix())
  for i := 0; i < 10; i++ {
    var intn int = rand.Intn(10)
    fmt.Println(intn)
  }
}
