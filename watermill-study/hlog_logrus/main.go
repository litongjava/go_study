/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
  "fmt"
  "github.com/sirupsen/logrus"
  "io"
  "log"
  "os"
  "path"
  "runtime"
  "time"

  "github.com/cloudwego/hertz/pkg/common/hlog"
  hertzlogrus "github.com/hertz-contrib/logger/logrus"
  "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
  //*hertzlogrus.Logger
  logger, done := getLogger()
  if done {
    return
  }

  //FullLogger
  hlog.SetLogger(logger)
  hlog.Info("Hi")

}

func getLogger() (*hertzlogrus.Logger, bool) {
  // Customizable output directory.
  var logFilePath string
  logFilePath = "logs/"
  if err := os.MkdirAll(logFilePath, 0o777); err != nil {
    log.Println(err.Error())
    return nil, true
  }

  // Set filename to date
  logFileName := time.Now().Format("2006-01-02") + ".log"
  fileName := path.Join(logFilePath, logFileName)
  if _, err := os.Stat(fileName); err != nil {
    if _, err := os.Create(fileName); err != nil {
      log.Println(err.Error())
      return nil, true
    }
  }

  // Initialize logger
  logger := hertzlogrus.NewLogger()
  lumberjackLogger := &lumberjack.Logger{
    Filename:   fileName,
    MaxSize:    20,   // Max file size 20M
    MaxBackups: 5,    // Keep up to 5 files
    MaxAge:     10,   // 10 days
    Compress:   true, // Compress log files
  }
  logOutput := io.MultiWriter(lumberjackLogger, os.Stdout)
  logger.SetOutput(logOutput)

  // Set log format
  formatter := &logrus.TextFormatter{
    FullTimestamp:   true,
    TimestampFormat: "2006/01/02 15:04:05.000000", // Microsecond precision
    CallerPrettyfier: func(f *runtime.Frame) (string, string) {
      return fmt.Sprintf("%s:%s:%d", f.File, f.Function, f.Line), ""
    },
    DisableQuote: true, // Disable quoting for cleaner output
  }
  logger.Logger().SetFormatter(formatter)
  logger.Logger().SetReportCaller(true)

  // Set debug level
  logger.SetLevel(hlog.LevelDebug)
  return logger, false
}
