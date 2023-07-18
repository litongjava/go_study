package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now()
	fmt.Println(timeNow.Format("2006-01-02 15:04:05")) // 2022-4-18 14:20:45
	fmt.Println(timeNow.Format("2006/01/02"))          // 2021/06/25

	// 获取当前时间戳
	fmt.Println(time.Now().Unix())
	// 指定的时间进行格式化
	fmt.Println(time.Unix(1650263305, 0).Format("2006-01-02 15:04:05"))

	//24小时之内的时间计算 ParseDuration
	//1小时1分1s之前
	t1, _ := time.ParseDuration("-1h1m1s")
	fmt.Println(t1)

	//24小时之外的时间计算
	beforeDay := timeNow.AddDate(0, 0, -1)   // 三个参数分别是年月日，此处获取的是前一天的日期
	beforeMonth := timeNow.AddDate(0, -1, 0) // 前一个月的日期
	beforeYear := timeNow.AddDate(-1, 0, 0)  // 去年的当天日期
	fmt.Println(beforeDay)
	fmt.Println(beforeMonth)
	fmt.Println(beforeYear)
	fmt.Println(beforeDay.Format("2006-01-02 15:04:05"))

	//24小时
	fmt.Println(timeNow.Format("2006-1-2 15:04:05.000 PM Mon Jan")) // 2021-6-25 10:59:05.410 AM Fri Jun
	fmt.Println(timeNow.Format("2006/01/02 15:04"))                 // 2021/06/25 10:59
	fmt.Println(timeNow.Format("2006-1-2 15:04:05.000"))            // 2021-6-25 10:59:05.410
	fmt.Println(timeNow.Format("Mon, 02 Jan 2006 15:04:05 GMT"))    // Fri, 25 Jun 2021 10:59:05 GMT
	// 12小时制
	fmt.Println(timeNow.Format("2006-01-02 03:04:05.000 PM Mon Jan")) // 2021-06-25 10:59:05.410 AM Fri Jun
	fmt.Println(timeNow.Format("15:04 2006/01/02"))                   // 10:59 2021/06/25
}
