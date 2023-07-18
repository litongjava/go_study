package main

import (
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "read-config-file/config"
)

func InitConfig() {
  yamlFile, err := ioutil.ReadFile("./config/config.yml")
  if err != nil {
    fmt.Println("error", err.Error())
  }
  var _config *config.Config
  err = yaml.Unmarshal(yamlFile, &_config)
  if err != nil {
    fmt.Println("error", err.Error())
  }
  fmt.Printf("config.app: %#v\n", _config.App)
  fmt.Printf("config.log: %#v\n", _config.Log)

}

func main() {
  InitConfig()
}
