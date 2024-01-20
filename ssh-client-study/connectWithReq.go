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
  log.Println("创建ssh channel 成功:", channel)

  go func() {
    for req := range inRequests {
      if req.WantReply {
        req.Reply(false, nil)
      }
    }
  }()

  modes := ssh.TerminalModes{
    ssh.ECHO:          1,
    ssh.TTY_OP_ISPEED: 14400,
    ssh.TTY_OP_OSPEED: 14400,
  }
  var modeList []byte
  for k, v := range modes {
    kv := struct {
      Key byte
      Val uint32
    }{k, v}
    modeList = append(modeList, ssh.Marshal(&kv)...)
  }

  modeList = append(modeList, 0)
  req := PtyRequestMsg{
    Term:     "xterm",
    Columns:  150,
    Rows:     32,
    Width:    uint32(150 * 8),
    Height:   uint32(32 * 8),
    Modelist: string(modeList),
  }
  ok, err := channel.SendRequest("pty-req", true, ssh.Marshal(&req))
  if !ok || err != nil {
    log.Println(err)
  }
  log.Println("send pty-req:", ok)
  ok, err = channel.SendRequest("shell", true, nil)
  if !ok || err != nil {
    log.Println(err)
  }

  log.Println("send shell:", ok)

  //发送命令
  cmd := "ls -l\n"
  _, err = channel.Write([]byte(cmd))
  if err != nil {
    return
  }
  log.Println("发送命令成功")

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
        log.Println("服务器响应:", string(buf[:n]))
      }
    }
  }()

  select {}
}
