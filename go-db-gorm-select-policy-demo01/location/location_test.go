package location

import (
  "fmt"
  "testing"
)

func TestLocation(t *testing.T) {
  location1 := location{lat: 1, long: 1}
  fmt.Println(location1)
  fmt.Println(&location1)
  fmt.Printf("%p\n", &location1)

  location2 := location1
  print(location1 == location2)
  var locationMap = map[location]string{}
  locationMap[location1] = "location1"
  locationMap[location2] = "location2"

  for key, value := range locationMap {
    fmt.Printf("%v %v \n", key, value)
  }
}
