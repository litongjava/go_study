package main

import (
  "encoding/base64"
  "encoding/json"
  "fmt"
  "github.com/gorilla/websocket"
  "log"
  "net/http"
  "strconv"
  "strings"
)

type ConnectionInfo struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Host     string `json:"host"`
  Port     int    `json:"port"`
}

func main() {
  fulldesc := "root@192.168.3.9"
  fmt.Println(fulldesc)

  fulldescArray := strings.Split(fulldesc, "@")
  username := fulldescArray[0]
  hostAndPort := strings.Split(fulldescArray[1], ":")
  host := hostAndPort[0]
  port := 22
  if len(hostAndPort) > 1 {
    port, _ = strconv.Atoi(hostAndPort[1])
  }

  connInfo := ConnectionInfo{
    Username: username,
    Password: "", // 将此处替换为实际密码
    Host:     host,
    Port:     port,
  }

  jsonBytes, _ := json.Marshal(connInfo)
  base64Str := base64.StdEncoding.EncodeToString(jsonBytes)

  // 启动WebSocket客户端并连接到WebSocket服务器
  var err error
  var websocketConn *websocket.Conn
  var response *http.Response
  //ws://localhost:5001/wsts/ws?msg={code}&rows=55&cols=116
  endPoint := "ws://localhost:5001/wsts/ws"
  urlStr := endPoint + "?msg=" + base64Str + "&rows=55&cols=116"
  websocketConn, response, err = websocket.DefaultDialer.Dial(urlStr, nil)
  if err != nil {
    log.Fatal("err:", err)
  } else {
    log.Println("response:", response)
  }
  defer websocketConn.Close()

  // 无限循环接收服务器的响应
  for {
    _, message, err := websocketConn.ReadMessage()
    if err != nil {
      log.Println("read:", err)
      return
    }
    log.Printf("recv: %s", message)
  }

  err = websocketConn.WriteMessage(websocket.TextMessage, []byte("pwd"))
  if err != nil {
    log.Println("write:", err)
    return
  }
  select {}
}
