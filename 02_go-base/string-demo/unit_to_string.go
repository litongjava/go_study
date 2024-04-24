package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myUint uint = 42
	myStr := strconv.FormatUint(uint64(myUint), 10) // 转换为 uint64, 然后转为字符串
	fmt.Println("The unsigned integer as a string is:", myStr)
}
