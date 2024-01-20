package main

import (
  "encoding/base64"
  "encoding/json"
  "golang.org/x/crypto/ssh"
  "io/ioutil"
  "log"
  "net"
  "strconv"
  "strings"
)

func init() {
  log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

type ConnectionInfo struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Host     string `json:"host"`
  Port     int    `json:"port"`
}

var private ssh.Signer

func main() {
  // read private key only once
  privateBytes, err := ioutil.ReadFile("ssh/ssh_host_rsa_key")
  if err != nil {
    log.Fatal("Failed to load private key: ", err)
  }

  private, err = ssh.ParsePrivateKey(privateBytes)
  if err != nil {
    log.Fatal("Failed to parse private key: ", err)
  }

  go sshServer()
  log.Println("listening...")
  select {}
}

func sshServer() {
  config := &ssh.ServerConfig{
    PasswordCallback: sshServerPasswordCallback,
  }

  config.AddHostKey(private)

  listener, err := net.Listen("tcp", "0.0.0.0:22")
  if err != nil {
    log.Fatalf("Failed to listen on 22 (%s)", err)
  }

  for {
    nConn, err := listener.Accept()
    if err != nil {
      log.Printf("Failed to accept incoming connection (%s)", err)
      continue
    }

    conn, chans, reqs, err := ssh.NewServerConn(nConn, config)
    if err != nil {
      log.Printf("Failed to handshake (%s)", err)
      continue
    }
    log.Printf("New SSH connection from %s (%s)", conn.RemoteAddr(), conn.ClientVersion())
    go ssh.DiscardRequests(reqs)

    for newChannel := range chans {
      // Each SSH session will have its own WebSocket connection
      jsonString := conn.Permissions.Extensions["json"]
      log.Println(jsonString)
      go handleChannel(newChannel)
    }
  }
}

func sshServerPasswordCallback(conn ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
  fulldesc := conn.User()
  password := string(pass)
  fulldescArray := strings.Split(fulldesc, "@")
  username := fulldescArray[0]
  hostAndPort := strings.Split(fulldescArray[1], ":")
  host := hostAndPort[0]
  var port int
  if len(hostAndPort) > 1 {
    port, _ = strconv.Atoi(hostAndPort[1])
  } else {
    port = 22
  }

  log.Println(host, port, username, password)

  connInfo := ConnectionInfo{
    Username: username,
    Password: password,
    Host:     host,
    Port:     port,
  }

  jsonBytes, _ := json.Marshal(connInfo)
  base64Str := base64.StdEncoding.EncodeToString(jsonBytes)

  return &ssh.Permissions{
    Extensions: map[string]string{
      "json": base64Str,
    },
  }, nil // 密码验证成功
}

func handleChannel(newChannel ssh.NewChannel) {
  if newChannel.ChannelType() != "session" {
    newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
    return
  }

  ch, reqs, err := newChannel.Accept()
  if err != nil {
    log.Println("could not accept channel.")
    return
  }
  defer ch.Close()

  // 发送'OK'到SSH客户端
  ch.Write([]byte("OK\r\n"))

  for req := range reqs {
    log.Println(req)
    ok := false
    switch req.Type {
    case "exec":
      ok = true
      log.Println("received command:", string(req.Payload))
      ch.Write([]byte("received\r\n"))
      req.Reply(ok, nil)
    case "pty-req":
      log.Println("received:", req.Type)
      ch.Write([]byte("received pty-req\r\n"))
      ok = true
      req.Reply(true, nil)
    case "shell":
      log.Println("received:", req.Type)
      ok = true
      req.Reply(true, nil)
      message := "# "
      ch.Write([]byte(message))
      go func() {
        buf := make([]byte, 256)
        for {
          n, err := ch.Read(buf)
          if err != nil {
            log.Println("读取输入失败:", err)
            break
          }
          if n > 0 {
            // 检查输入是否为回车
            if buf[0] == '\r' {
              // 如果是回车，则输出"#"
              _, err = ch.Write([]byte("\n\r success \n\r# "))
              if err != nil {
                log.Println("写入输出失败:", err)
                break
              }
            } else {
              // 如果不是回车，则回显输入到客户端
              log.Println("shell input:", string(buf[:n]))
              _, err = ch.Write(buf[:n])
              if err != nil {
                log.Println("写入输出失败:", err)
                break
              }
            }
          }
        }
      }()
    }
    if !ok {
      message := "declining request:" + req.Type + "\r\n"
      log.Println(message)
      ch.Write([]byte(message))
      req.Reply(ok, nil)
    }
  }
}
