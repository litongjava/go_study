package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	marshalJson()
}

type CmdResult struct {
	Success bool   `json:"success"`
	Output  string `json:"output"`
	Time    int64  `json:"time"`
}

func marshalJson() {
	c := CmdResult{}
	c.Success = true
	c.Output = "litong"
	c.Time = 1
	fmt.Println(c)

	bytes, _ := json.Marshal(c)
	fmt.Println(string(bytes))

	bytes1, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println(string(bytes1))
}
