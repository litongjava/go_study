package main

import (
  "bytes"
  "fmt"
  "golang.org/x/crypto/ssh"
  "log"
  "time"
)

func main() {
  sshHost := "192.168.3.9"
  sshUser := "root"
  sshPasswrod := "Cttic@2013"
  sshType := "password" // password或者key
  //sshKeyPath := "" // ssh id_rsa.id路径
  sshPort := 22

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
  defer sshClient.Close()

  // 创建ssh-session
  session, err := sshClient.NewSession()
  if err != nil {
    log.Fatal("创建ssh session失败", err)
  }

  defer session.Close()

  // 请求一个伪终端
  if err := session.RequestPty("xterm", 80, 40, ssh.TerminalModes{}); err != nil {
    log.Println(err)
  }

  // 创建一个缓冲区来保存会话的输出
  var b bytes.Buffer
  session.Stdout = &b

  // 启动会话。这应该会自动发送欢迎消息
  if err := session.Shell(); err != nil {
    log.Println(err)
  }

  // 等待会话结束。注意这可能需要一些时间，你可能需要考虑使用timeout
  if err := session.Wait(); err != nil {
    log.Println(err)
  }

  // 打印欢迎消息
  fmt.Println(b.String())

  // 执行远程命令
  combo, err := session.CombinedOutput("whoami; cd /; ls -al;")
  if err != nil {
    log.Fatal("远程执行cmd失败", err)
  }
  log.Println("命令输出:", string(combo))
}

//func publicKeyAuthFunc(kPath string) ssh.AuthMethod  {
//	keyPath ,err := homedir.Expand(kPath)
//	if err != nil {
//		log.Fatal("find key's home dir failed",err)
//	}
//
//	key,err := ioutil.ReadFile(keyPath)
//	if err != nil {
//		log.Fatal("ssh key file read failed",err)
//	}
//
//	signer,err := ssh.ParsePrivateKey(key)
//	if err != nil {
//		log.Fatal("ssh key signer failed",err)
//	}
//	return ssh.PublicKeys(signer)
//}
