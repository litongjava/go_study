package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  //v1
  //var s, sep string

  //for i := 1; i < len(os.Args); i++ {
  //  s += sep + os.Args[i]
  //  sep = " "
  //}

  //v2
  //s, sep := "", ""
  //for _, arg := range os.Args[1:] {
  //  s += sep + arg
  //  sep = " "
  //}

  //v3

  fmt.Println(strings.Join(os.Args[1:], " "))
}
