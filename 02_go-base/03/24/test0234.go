package main

import (
  "fmt"
  "sort"
)

func test032402() {
  var nowhere *int
  if nowhere != nil {
    fmt.Println(*nowhere)
  }
  fmt.Println(nowhere)
}

type person struct {
  age int
}

func (p *person) birthday() {
  if p != nil {
    p.age++
  }
}
func test032404() {
  var nobody *person
  fmt.Println(nobody)

  nobody.birthday()
}
func test032406() {
  var fn func(a, b int) int
  fmt.Println(fn == nil)
}

func sortStrings(s []string, less func(i, j int) bool) {
  if less == nil {
    less = func(i, j int) bool {
      return s[i] < s[j]
    }
  }
  sort.Slice(s, less)
}
func test03240602() {
  food := []string{"onion", "carrot", "celery"}
  sortStrings(food, nil)
  fmt.Println(food)
}
func test032408() {
  var soup []string
  fmt.Println(soup == nil)

  for _, ingredient := range soup {
    fmt.Println(ingredient)
  }

  fmt.Println(len(soup))

  soup = append(soup, "onion", "carrot", "celery")
  fmt.Println(soup)
}
func mirepoix(ingredients []string) []string {
  return append(ingredients, "onion", "carrot", "celery")
}
func test03240802() {
  soup := mirepoix(nil)
  fmt.Println(soup)
}
func test032410() {
  var soup map[string]int
  fmt.Println(soup == nil)

  //对值为nil的map进行读取,不会报错,但是ok的值为false
  measurement, ok := soup["onion"]
  if ok {
    fmt.Println(measurement)
  }

  for ingredient, measurement := range soup {
    fmt.Println(ingredient, measurement)
  }
}
func test032412() {
  var v interface{}
  fmt.Printf("%T %v %v\n", v, v, v == nil)

  var p *int
  //此时v的类型是指向int类型的指针
  v = p
  fmt.Printf("%T %v %v\n", v, v, v == nil)
  //检验接口变量的内部表示
  fmt.Printf("%#v\n", v)
}

type number struct {
  value int
  valid bool
}

func newNumber(v int) number {
  return number{value: v, valid: true}
}
func (n number) String() string {
  if !n.valid {
    return "not set"
  }
  return fmt.Sprintf("%d", n.value)
}
func test032413() {
  n := newNumber(42)
  fmt.Println(n)

  e := number{}
  fmt.Println(e)
}
func main() {
  test032413()
}
