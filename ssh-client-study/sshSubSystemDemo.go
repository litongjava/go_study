package main

import (
  "fmt"
  "golang.org/x/crypto/ssh"
  "log"
  "time"
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

  // 创建ssh-session
  session, err := sshClient.NewSession()
  if err != nil {
    log.Fatal("创建ssh session 失败", err)
  }
  log.Println("创建ssh session 成功:", session)

  channel, inRequests, err := sshClient.OpenChannel("session", nil)
  if err != nil {
    log.Fatal("创建ssh channel 失败", err)

  }
  log.Println("创建ssh channel 成功:", channel, inRequests)

  go func() {
    for req := range inRequests {
      if req.WantReply {
        req.Reply(false, nil)
      }
    }
  }()

  // Send subsystem request.
  subsystemPayload := struct {
    Subsystem string
  }{
    Subsystem: "sftp", // or other subsystem you want to start
  }

  ok, err := channel.SendRequest("subsystem", true, ssh.Marshal(&subsystemPayload))
  if !ok || err != nil {
    log.Println(err)
  }
  log.Println("send subsystem request:", ok)

  //第二个协程将远程主机的返回结果返回给用户
  go func() {
    buf := make([]byte, 1024)
    for {
      n, err := channel.Read(buf)
      if err != nil {
        log.Println("读取服务器响应失败", err)
        return
      }
      if n > 0 {
        log.Println("redceived from ssh server:", string(buf[:n]))
      }
    }
  }()

  select {}
}
