package main

import "fmt"

type kelvin float64

//sensor funciton type
type sensor func() kelvin

func realSensor() kelvin {
  return 0
}

func calibrate(s sensor, offset kelvin) sensor {
  //返回一返回值是kelvin的函数
  return func() kelvin {
    return s() + offset
  }
}
func main() {
  //sensor的类型是sensor func(),读作一等函数sensor
  var offset kelvin = 5
  sensor := calibrate(realSensor, offset)
  fmt.Println(sensor()) //->5
  offset = 6
  fmt.Println(sensor()) //->5
  sensor = calibrate(realSensor, offset)
  fmt.Println(sensor()) //->6
}
