package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myUint uint = 42
	myUintPtr := &myUint // 指向 myUint 的指针

	// 检查指针是否为 nil，然后解引用并转换
	if myUintPtr != nil {
		myStr := strconv.FormatUint(uint64(*myUintPtr), 10) // 解引用指针并转为 uint64, 然后转为字符串
		fmt.Println("The unsigned integer as a string is:", myStr)
	} else {
		fmt.Println("Pointer is nil")
	}
}
