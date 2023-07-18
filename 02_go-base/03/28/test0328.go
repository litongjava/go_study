package main

import (
  "fmt"
  "image"
  "log"
  "sync"
  "time"
)

var mu sync.Mutex

func test03280201() {
  mu.Lock()
  defer mu.Unlock()
}

type Visited struct {
  mu      sync.Mutex
  visited map[string]int
}

func (v *Visited) VisitLink(url string) int {
  v.mu.Lock()
  defer v.mu.Unlock()
  count := v.visited[url]
  count++
  v.visited[url] = count
  return count
}

func worker() {
  pos := image.Point{X: 10, Y: 10}
  direction := image.Point{X: 1, Y: 0}
  next := time.After(time.Second)
  for {
    select {
    case <-next:
      pos = pos.Add(direction)
      fmt.Println("current position is", pos)
      next = time.After(time.Second)
    }
  }
}

type command int

const (
  right = command(0)
  left  = command(1)
)

type RoverDriver struct {
  commandc chan command
}

func NewRoverDriver() *RoverDriver {
  r := &RoverDriver{
    commandc: make(chan command),
  }
  go r.drive()
  return r
}
func (r *RoverDriver) drive() {
  pos := image.Point{X: 0, Y: 0}
  direction := image.Point{X: 1, Y: 0}
  updateInterval := 250 * time.Microsecond
  nextMove := time.After(updateInterval)

  for {
    select {
    case c := <-r.commandc:
      switch c {
      case right:
        direction = image.Point{
          X: -direction.Y,
          Y: direction.X,
        }
      case left:
        direction = image.Point{
          X: direction.X,
          Y: -direction.X,
        }
      }
      log.Printf("new direction %v", direction)
    case <-nextMove:
      pos = pos.Add(direction)
      log.Printf("moved to %v", pos)
      nextMove = time.After(updateInterval)
    }
  }
}

func (r *RoverDriver) Left() {
  r.commandc <- left
}

func (r *RoverDriver) Right() {
  r.commandc <- right
}
func test032807() {
  r := NewRoverDriver()
  time.Sleep(3 * time.Second)
  r.Left()
  time.Sleep(3 * time.Second)
  r.Right()
  time.Sleep(3 * time.Second)
}
func main() {
  test032807()
}
