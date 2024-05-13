package tasks

import (
  "bytes"
  "fmt"
  "github.com/sirupsen/logrus"
  "io"
  "log"
  "os"
  "path"
  "strings"
  "time"
)

var llog *logrus.Logger

type CustomFormatter struct {
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
  b := &bytes.Buffer{}

  timestamp := entry.Time.Format("2006-01-02 15:04:05.000000")
  level := strings.ToUpper(entry.Level.String())
  fileName := path.Base(entry.Caller.File)

  fs := strings.Split(entry.Caller.Function, ".")
  funName := ""
  if len(fs) > 0 {
    funName = fs[len(fs)-1]
  }

  fmt.Fprintf(b, "%s %s %s:%d %s: %s", level, timestamp, fileName, entry.Caller.Line, funName, entry.Message)

  for key, value := range entry.Data {
    fmt.Fprintf(b, " %s=%+v", key, value)
  }

  b.WriteByte('\n')
  return b.Bytes(), nil
}

func init() {
  if llog != nil {
    return
  }
  // 实例化
  innerLogger := logrus.New()
  output, err := createOutputFile()
  // 设置输出
  if err != nil {
    innerLogger.Out = os.Stdout
  } else {
    innerLogger.Out = io.MultiWriter(output, os.Stdout)
  }

  // 设置日志级别
  innerLogger.SetLevel(logrus.InfoLevel)
  //初始化日志
  innerLogger.SetReportCaller(true) //需要设置这个为true
  innerLogger.SetFormatter(&CustomFormatter{})

  llog = innerLogger
}

func createOutputFile() (*os.File, error) {
  now := time.Now()
  logFilePath := ""
  if dir, err := os.Getwd(); err == nil {
    logFilePath = dir + "/logs/" // 日志文件存储路径
  }
  _, err := os.Stat(logFilePath)
  if os.IsNotExist(err) {
    if err := os.MkdirAll(logFilePath, 0777); err != nil {
      log.Println(err.Error())
      return nil, err
    }
  }
  logFileName := now.Format("2006-01-02") + ".log"
  // 日志文件
  fileName := path.Join(logFilePath, logFileName)
  if _, err := os.Stat(fileName); err != nil {
    if _, err := os.Create(fileName); err != nil {
      return nil, err
    }
  }
  // 写入文件
  src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
  if err != nil {
    return nil, err
  }
  return src, nil
}
