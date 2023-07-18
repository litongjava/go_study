package main

import (
  "errors"
  "fmt"
  "io"
  "io/ioutil"
  "os"
  "strings"
)

func test032501() {
  files, err := ioutil.ReadDir(".")
  //files, err := ioutil.ReadDir("unicorns")
  //files, err := ioutil.ReadDir("C:\\Program Files\\Git\\etc\\hosts")
  if err != nil {
    fmt.Println(err)
    //传入非0值,表示程序发生错误,退出
    os.Exit(1)
  }

  for _, file := range files {
    fmt.Println(file.Name())
  }
}

func proverbs(name string) error {
  f, err := os.Create(name)
  if err != nil {
    return err
  }
  defer f.Close()
  _, err = fmt.Fprintln(f, "Errors ar values")

  if err != nil {
    return err
  }

  _, err = fmt.Fprintln(f, "Don't just check errors,handle them gracefully")
  return err
}
func test032505() {
  err := proverbs("proverbs.txt")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

type safeWriter struct {
  w   io.Writer
  err error
}

func (sw *safeWriter) writeln(s string) {
  if sw.err != nil {
    return
  }
  _, sw.err = fmt.Fprintln(sw.w, s)
}
func proverbsv2(name string) error {
  f, err := os.Create(name)
  if err != nil {
    return err
  }
  defer f.Close()

  sw := safeWriter{w: f}
  sw.writeln("Erros are values")
  sw.writeln("Don't just check erros,handle them gracefully")
  sw.writeln("Don't panic")

  return sw.err
}
func test032510() {

}

const rows, columns = 9, 9

type Grid [rows][columns]int8

var (
  ErrBounds = errors.New("out of bounds")
  ErrDigit  = errors.New("invalid digit")
)

type SudokuError []error

func (se SudokuError) Error() string {
  var s []string
  for _, err := range se {
    s = append(s, err.Error())
  }
  return strings.Join(s, ", ")
}
func (g *Grid) Set(row, column int, digit int8) error {
  var errs SudokuError
  if !inBounds(row, column) {
    errs = append(errs, ErrBounds)
  }
  if !validDigit(digit) {
    errs = append(errs, ErrDigit)
  }

  if len(errs) > 0 {
    return errs
  }
  g[row][column] = digit
  return nil
}
func validDigit(digit int8) bool {
  return false
}
func inBounds(row, column int) bool {
  if row < 0 || row <= rows {
    return false
  }
  if column < 0 || column <= columns {
    return false
  }
  return true
}
func test032517() {
  var g Grid
  err := g.Set(12, 0, 15)
  if err != nil {
    switch err {
    case ErrBounds, ErrDigit:
      fmt.Println("Errors")
    default:
      fmt.Println(err)
    }
    os.Exit(1)
  }
}
func test032512() {
  var g Grid
  err := g.Set(10, 0, 5)
  if err != nil {
    fmt.Printf("An error occurred: %v.\n", err)
    os.Exit(1)
  }
}
func test032515() {
  var g Grid
  err := g.Set(0, 0, 15)
  if err != nil {
    switch err {
    case ErrBounds, ErrDigit:
      fmt.Println("Les crreurs de parameters hors limites")
    default:
      fmt.Println(err)
    }
    os.Exit(1)
  }
}
func test032519() {
  var g Grid
  err := g.Set(10, 0, 15)
  if err != nil {
    //判断err的类型是否为ShduKuError,如果是ok为true
    if errs, ok := err.(SudokuError); ok {
      fmt.Printf("%d error(s) occurred:\n", len(errs))

      for _, e := range errs {
        fmt.Printf("_ %v\n", e)
      }
    }
    os.Exit(1)
  }
}
func test032525() {
  var zero int
  _ = 42 / zero
}
func test032527() {
  //先执行panic在执行defer
  defer func() {
    //使用recover获取panic的传入值
    if e := recover(); e != nil {
      fmt.Println("recover")
      fmt.Println(e)
    }
  }()

  panic("I forget my towel")
}
func main() {
  //test032519()
  test032527()
}
