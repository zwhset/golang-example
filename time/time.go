package main

import (
	"fmt"
	"time"
)

func main() {
	// Format time now
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))

	// String to time
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	fmt.Println(t1)

	// format
	fmt.Println(t.Format("3:04PM"))
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t.Format("2006-01-02T15:04:05.999999-07:00"))
	fmt.Println(t.Format("3 04 PM"))
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	// 时间戳
	fmt.Println()
	epoch()

	// 时间方法
	fmt.Println()
	eTime()

	// OUTPUT
//	2018-09-18T21:49:44+08:00
//	2012-11-01 22:08:41 +0000 +0000
//	9:49PM
//	Tue Sep 18 21:49:44 2018
//	2018-09-18T21:49:44.759508+08:00
//	9 49 PM
//	2018-09-18 21:49:44
//
//now:  2018-09-18 21:49:44.759696872 +0800 CST m=+0.000634545
//secs:  1537278584
//nanos:  1537278584759696872
//	millis 15372785847596
//	2018-09-18 21:49:44 +0800 CST
//	2018-09-18 21:49:44.759696872 +0800 CST
//
//	2018-09-18 21:49:44.759719252 +0800 CST m=+0.000656926
//	2009-11-17 20:34:58.651387237 +0000 UTC
//	2009 November 17 20
//	2018 September 18
//	true
//	false
//	false
//	77441h14m46.108332015s
//	77441.24614120334
//	4.6464747684722e+06
//	2.7878848610833204e+08
//	2018-09-18 13:49:44.759719252 +0000 UTC
//	true
//	2001-01-17 03:20:12.543055222 +0000 UTC
}

func epoch()  {

	now := time.Now() // now
	secs := now.Unix() // 时间戳
	nanos := now.UnixNano() // 纳米秒的时间戳
	fmt.Println("now: ", now)
	fmt.Println("secs: ", secs)
	fmt.Println("nanos: ", nanos)

	millis := nanos / 100000 // 毫秒的时间戳
	fmt.Println("millis", millis)

	fmt.Println(time.Unix(secs, 0)) // 秒的时间戳转时间
	fmt.Println(time.Unix(0, nanos)) // 纳米秒的时间戳转时间
}

func eTime()  {
	now := time.Now()
	fmt.Println(now)

	// 生成时间对象
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(then)

	// 访问时间属性
	fmt.Println(then.Year(), then.Month(), then.Day(), then.Hour()) //...
	fmt.Println(now.Year(), now.Month(), now.Day()) // ...

	// 时间判断
	fmt.Println(then.Before(now)) // 之前
	fmt.Println(then.After(now)) // 之后
	fmt.Println(then.Equal(now)) // 相等

	diff := now.Sub(then) // 时间差
	fmt.Println(diff)
	fmt.Println(diff.Hours()) // 差多少小时
	fmt.Println(diff.Minutes()) // 差多少分钟
	fmt.Println(diff.Seconds()) // 差多少秒

	// 添加时间
	fmt.Println(then.Add(diff)) // 增加时间
	fmt.Println(then.Add(diff).Equal(now)) // 增加时间差两者时间相等
	fmt.Println(then.Add(-diff)) // 减少时间
}