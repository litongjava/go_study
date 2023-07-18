package main

import (
  "log"
  "os/exec"
  "time"
)

func main() {
  //命令执行失败：exec: "xyz": executable file not found in %PATH%
  cmdResult := runCmdbyGrep("ping", "www.baidu.com")
  log.Println(cmdResult)

}

type CmdResult struct {
  success bool   `json:"success"`
  output  string `json:"output"`
  time    int64  `json:"time"`
}

func runCmdbyGrep(name string, arg ...string) CmdResult {
  start := time.Now().Unix()
  command := exec.Command(name, arg...)
  //会自动执行命令
  result, err := command.CombinedOutput()
  end := time.Now().Unix()

  cmdResult := CmdResult{}
  cmdResult.time = end - start
  if err != nil {
    cmdResult.success = false
    cmdResult.output = err.Error()
  } else {
    cmdResult.success = true
    cmdResult.output = (string(result))
  }
  return cmdResult
}
