package main

import (
  "fmt"
  "math/rand"
  "strings"
  "time"
)

func sleepyGoPher(id int) {
  time.Sleep(3 * time.Second)
  fmt.Println("...snore...", id)
}

func test032704() {
  for i := 0; i < 5; i++ {
    go sleepyGoPher(i)
  }
  time.Sleep(4 * time.Second)
}
func test032703() {
  go sleepyGoPher(0)          //分支线路
  time.Sleep(4 * time.Second) //主线路
  //main函数执行完,所有的go routine都会停止
}
func sleepyGoPherForChannel(id int, c chan int) {
  time.Sleep(3 * time.Second)
  fmt.Println("...snore...", id)
  c <- id
}
func test032708() {
  c := make(chan int)
  for i := 0; i < 5; i++ {
    go sleepyGoPherForChannel(i, c)
  }
  for i := 0; i < 5; i++ {
    gopherId := <-c //等待
    fmt.Println("gopher:", gopherId, "has finished sleeping")
  }
}

func sleepyGoPherForSelect(id int, c chan int) {
  duration := time.Duration(rand.Intn(4000))
  fmt.Println(duration)
  //0-4s之间的随机时间
  time.Sleep(duration * time.Microsecond)
  c <- id
}
func test032710() {
  c := make(chan int)
  for i := 0; i < 5; i++ {
    go sleepyGoPherForSelect(i, c)
  }

  //创建一个通道,超时时间是2s
  timeout := time.After(2 * time.Second)
  for i := 0; i < 5; i++ {
    select {
    //从通道中接收值,如果2s之后没有接收到值则执行下一个通知
    case gopherId := <-c:
      fmt.Println("gopher", gopherId, " has finished sleeping")
    case <-timeout:
      fmt.Println("my patience ran out")
      return
    }
  }
}
func test032713() {
  c := make(chan int)
  go func() { c <- 2 }()
  i := <-c
  fmt.Println(i)
}
func sourceGopher(downstream chan string) {
  for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
    downstream <- v
  }
  //downstream <- ""
  close(downstream)
}

//func filterGopher(upstream, downstream chan string) {
//  for {
//    //item := <-upstream
//    //ok为false表示通道已关闭
//    item, ok := <-upstream
//    if !ok {
//      close(downstream)
//      return
//    }
//    if !strings.Contains(item, "bad") {
//      downstream <- item
//    }
//  }
//}
func filterGopher(upstream, downstream chan string) {
  for item := range upstream {
    if !strings.Contains(item, "bad") {
      downstream <- item
    }
  }
  close(downstream)
}

//func printGopher(upstream chan string) {
//  for {
//    v := <-upstream
//    if v == "" {
//      return
//    }
//    fmt.Println(v)
//  }
//}
func printGopher(upstream chan string) {
  for v := range upstream {
    fmt.Println(v)
  }
}
func test032715() {
  c0 := make(chan string)
  c1 := make(chan string)
  go sourceGopher(c0)
  go filterGopher(c0, c1)
  printGopher(c1)
}
func main() {
  test032715()
}
