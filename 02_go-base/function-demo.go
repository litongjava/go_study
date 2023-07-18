package main

import "fmt"

func main() {
  fmt.Println("function-demo")
  func() {
    fmt.Println("this is a anonymous function")
  }()
}
