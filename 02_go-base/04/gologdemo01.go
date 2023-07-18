package main

import (
  "log"
  "os"
)

func init() {
  //设置前缀
  log.SetPrefix("Demo ")
  //如果文件不存在则创建,如果文件存在在追加,文件是只读的,并指定文件权限为0666
  f, err := os.OpenFile("demo.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
  if err != nil {
    log.Fatalln(err)
  }
  //设置输出到文件
  log.SetOutput(f)
  //设置Flats为 日期 时间 微秒 文件名:行号
  log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}
func main() {
  log.Println("Helle")
}
