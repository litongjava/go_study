package main

import (
  "fmt"
)

//type kelvin float64

func main() {
  var k kelvin = 294.0
  sensor := func() kelvin {
    return k
  }
  fmt.Println(sensor())
  k++
  fmt.Println(sensor())
}
