package main

import (
  "flag"
  "fmt"
)

func main() {
  boolVal := flag.Bool("testBool", false, "testBool is bool type.")
  flag.Parse()

  // 如果使用 -testBool作为参数，控制台将会打印 true, 否则打印 false
  fmt.Println(*boolVal)
}
