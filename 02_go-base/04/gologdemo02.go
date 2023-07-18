package main

import (
  "io"
  "io/ioutil"
  "log"
  "os"
)

var (
  Trace   *log.Logger //几乎任何东西
  Debug   *log.Logger //调试
  Info    *log.Logger //重要信息
  Warning *log.Logger //警告
  Error   *log.Logger //错误
)

func init() {
  file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
  if err != nil {
    log.Fatalln("can't open log file", err)
  }
  Trace = log.New(ioutil.Discard, "TRACE ", log.Ldate|log.Ltime|log.Lshortfile)
  Debug = log.New(os.Stdout, "Debug ", log.Ldate|log.Ltime|log.Lshortfile)
  Info = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
  Warning = log.New(os.Stdout, "WARNING ", log.Ldate|log.Ltime|log.Lshortfile)
  Error = log.New(io.MultiWriter(file, os.Stdout), "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
  Trace.Println("鸡毛蒜皮的小时")
  Debug.Println("调试信息")
  Info.Println("特别的信息")
  Warning.Println("警告")
  Error.Println("出现了故障")
}
