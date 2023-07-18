package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "nginx -s reload"
	//拆分,返回切片
	stringSlice := strings.Split(text, " ")
	//获取第0个元素
	fmt.Println(stringSlice[0])
	//获取从1到最后一个元素
	fmt.Println(stringSlice[1:])
}
