package main

import (
  "fmt"
  "log"
  "time"

  "golang.org/x/crypto/ssh"
)

func init() {
  log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

func main() {
  sshHost := "192.168.3.9"
  sshPort := 22
  sshUser := "root"
  sshPasswrod := "Cttic@2013"
  sshType := "password" // password或者key
  //sshKeyPath := "" // ssh id_rsa.id路径

  // 创建ssh登录配置
  config := &ssh.ClientConfig{
    Timeout:         time.Second, // ssh连接time out时间一秒钟,如果ssh验证错误会在一秒钟返回
    User:            sshUser,
    HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 这个可以,但是不够安全
    //HostKeyCallback: hostKeyCallBackFunc(h.Host),
  }
  if sshType == "password" {
    config.Auth = []ssh.AuthMethod{ssh.Password(sshPasswrod)}
  } else {
    //config.Auth = []ssh.AuthMethod(publicKeyAuthFunc(sshKeyPath))
    return
  }

  // dial 获取ssh client
  addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
  sshClient, err := ssh.Dial("tcp", addr, config)
  if err != nil {
    log.Fatal("创建ssh client 失败", err)
  }
  log.Println("创建ssh client 成功:", sshClient)
  defer sshClient.Close()

  // 创建ssh-session
  session, err := sshClient.NewSession()
  if err != nil {
    log.Fatal("创建ssh session失败", err)
  }
  log.Println("创建ssh session成功:", session)
  defer session.Close()

  // 运行远程命令
  cmd := "ls -l"
  output, err := session.Output(cmd)
  if err != nil {
    log.Fatal("运行远程命令失败", err)
  }
  fmt.Println(string(output))

}
