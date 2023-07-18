package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  fmt.Println("What's your name?")
  //从stdin中读取内容
  reader := bufio.NewReader(os.Stdin)
  //使用换行作为输入的结束,换行符也会读入到text中
  text, _ := reader.ReadString('\n')

  fmt.Printf("Your name is: %s", text)
}
