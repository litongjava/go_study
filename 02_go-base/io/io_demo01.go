package main

import (
  "fmt"
  "io"
  "os"
)

func main() {
  // 创建一个 io.Pipe
  //var pw *io.PipeWriter
  //var pr *io.PipeReader
  pr, pw := io.Pipe()

  // 启动一个 goroutine 写入数据
  go func() {
    defer pw.Close()
    _, err := pw.Write([]byte("Hello, Pipe!\n"))
    if err != nil {
      fmt.Println("Error writing to pipe:", err)
    }
  }()

  // 读取数据并打印
  _, err := io.Copy(os.Stdout, pr)
  if err != nil {
    fmt.Println("Error reading from pipe:", err)
  }
}
