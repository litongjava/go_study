package main

import (
  "fmt"
  "strings"
  "time"
)

func test032202() {
  answer := 42
  address := &answer
  fmt.Println(address)
  fmt.Println(*address)
}

func test03220501() {
  answer := 42
  address := &answer //*int
  fmt.Printf("address is a %T\n", address)
}
func test03220502() {
  canada := "Canada"
  var home *string
  fmt.Printf("Home is a %T\n", home)

  home = &canada
  fmt.Printf(*home)
}
func test032307() {
  var administrator *string
  scolese := "Christopher J. Scolese"
  administrator = &scolese
  fmt.Println(*administrator)

  bolden := "Charles F. Bolden"
  administrator = &bolden
  fmt.Println(administrator)
  fmt.Println(*administrator)

  bolden = "Charles Frank Bolden Jr"
  //变量的地址值没有改变,变量的内容值改变了,所以取内容值取到的内容也会改变
  fmt.Println(administrator)
  fmt.Println(*administrator)

  //通过内容值改变内容值
  *administrator = "Maj. Gen. Charles Frank Bolden Jr."
  fmt.Println(bolden)

  //将新的指针类型的变量指向地址值
  major := administrator
  *major = "Major General Charles Frank Bolden Jr."
  fmt.Println(bolden)

  //比较内存地址是否相等
  fmt.Println(administrator == major)

  lightfoot := "Robert M. Lightfoot Jr."
  administrator = &lightfoot
  fmt.Println(administrator == major)

  //将内存地址产生一个副本赋值给这个变量
  charles := *major
  *major = "Charles Bolden"
  fmt.Println(charles)
  fmt.Println(bolden)

  //字符串值相等,
  charles = "Charles Bolden"
  fmt.Println(charles == bolden)
  fmt.Println(&charles == &bolden)
}
func test032309() {
  type person struct {
    name, superpower string
    age              int
  }
  timmy := &person{
    name: "Timothy",
    age:  10,
  }

  (*timmy).superpower = "Flying"
  timmy.superpower = "Flying"
  fmt.Printf("%+v\n", timmy)
}

func test032311() {
  //和结构体一样，可以把 & 放在数组的复合字面值前面来创建指向数组的指针。
  superpowers := &[3]string{"flight", "invisibility", "super strength"}
  fmt.Println(superpowers[0])
  fmt.Println(superpowers[1:2])
}

func birthday(p *person) {
  p.age++
}
func birthday2(p person) {
  p.age++
}
func test03231301() {
  rebecca := person{
    name:       "Rebecca",
    superpower: "imageination",
    age:        14,
  }
  birthday(&rebecca)
  birthday2(rebecca)
  fmt.Printf("%+v\n", rebecca)
}

type person struct {
  name, superpower string
  age              int
}

func (p *person) birthday() {
  p.age++
}
func test032315() {
  //Person的解引用
  terry := &person{
    name: "Terry",
    age:  15,
  }
  terry.birthday()
  fmt.Printf("%+v\n", terry)
  //Person
  nathan := person{
    name: "Nathan",
    age:  17,
  }
  //通过.进行调用,会自动通过&符号调用变量的内存地址
  nathan.birthday()
  (&nathan).birthday()
  fmt.Printf("%+v\n", nathan)

  const layout = "Mon,Jan 2,2006"
  day := time.Now()
  tomorrow := day.Add(24 * time.Hour)

  fmt.Println(day.Format(layout))
  fmt.Println(tomorrow.Format(layout))
}

type stats struct {
  level             int
  endurance, health int
}

func levelUp(s *stats) {
  s.level++
  s.endurance = 42 + (14 * s.level)
  s.health = 5 * s.endurance
}

type character struct {
  name  string
  stats stats
}

func test032317() {
  player := character{name: "Matthias"}
  //使用&解地址符号,获取结构体内部的字段的地址值
  levelUp(&player.stats)

  fmt.Printf("%+v\n", player.stats)
}
func reset(board *[8][8]rune) {
  board[0][0] = 'r'
}
func test032318() {
  var board [8][8]rune
  reset(&board)

  fmt.Printf("%c", board[0][0])
}

func reclassify(planets *[]string) {
  *planets = (*planets)[0:8]
}
func test032324() {
  planets := []string{
    "Mercury", "Venus", "Earth", "Mars",
    "Jupiter", "Saturn", "Uranus", "Neptune",
    "Pluto",
  }
  reclassify(&planets)
  fmt.Println(planets)
}

type talker interface {
  talk() string
}

func shout(t talker) {
  louder := strings.ToUpper(t.talk())
  fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
  return "nack nack"
}

type laser int

func (l *laser) talk() string {
  return strings.Repeat("pew ", int(*l))
}
func test032325() {
  shout(martian{})
  shout(&martian{})

  pew := laser(2)
  shout(&pew)
}
func main() {
  test032325()
}
