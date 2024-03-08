package main

import (
	"fmt"
	"github.com/bytedance/sonic"
)

func main() {
	data := map[string]interface{}{}
	data["ContentType"] = "audio"
	data["Audio"] = []byte{0x00, 0x01}
	byteData, err := sonic.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byteData))
}
