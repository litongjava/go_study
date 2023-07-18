package main

import (
  "fmt"
  "strings"
  "time"
)

type talkder interface {
  talk() string
}

type martian struct{}

func (m martian) talk() string {
  return "nack nack"
}

type laser int

func (l laser) talk() string {
  return strings.Repeat("pew ", int(l))
}
func shout(t talkder) {
  louder := strings.ToUpper(t.talk())
  fmt.Println(louder)
}

//接口可以和struct嵌入一起使用
type startship struct {
  laser
}

func test032202() {
  shout(martian{})
  shout(laser((2)))

  s := startship{laser(3)}
  fmt.Println(s.talk())
  shout(s)
}

type stardater interface {
  YearDay() int
  Hour() int
}
type sol int

func (s sol) YearDay() int {
  return int(s % 668)
}

func (s sol) Hour() int {
  return 0
}

//stardate returns a ficitional measure of time
func stardate(t stardater) float64 {
  doy := float64(t.YearDay())
  h := float64(t.Hour()) / 24.0
  return 1000 + doy + h
}
func test032203() {
  day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
  fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

  s := sol(1422)
  fmt.Printf("%1.f Happy birthday\n", stardate(s))
}

type location struct {
  lat, long float64
}

func (l location) String() string {
  return fmt.Sprintf("%v,%v", l.lat, l.long)
}

type coordinate struct {
  d, m, s float64
  h       rune
}
func (c coordinate) String() string{
  return fmt.Sprintf("Elysium Planitia is at %v°$v N, %v $v E")
}
func test032207() {

}
func test032206() {
  curiosity := location{-4.5896, 127.4417}
  fmt.Println(curiosity)
}
func main() {
  test032207()
}
