package main

import (
  "fmt"
  "sync"
)

type Singleton struct{}

var (
  instance *Singleton
  once     sync.Once
)

func GetInstance() *Singleton {
  once.Do(func() {
    instance = &Singleton{}
  })
  return instance
}

func main() {
  //var wg sync.WaitGroup
  //
  //for i := 0; i < 5; i++ {
  //  wg.Add(1)
  //  go func() {
  //    defer wg.Done()
  //    s := GetInstance()
  //    fmt.Printf("Singleton instance address: %p\n", s)
  //  }()
  //}
  //
  //wg.Wait()
  s := GetInstance()
  fmt.Printf("Singleton instance address: %p\n", s)
}
