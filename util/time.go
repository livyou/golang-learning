package util

import (
	"fmt"
	"time"
)

//当前时间戳
func TimeStamp() {
	now := time.Now()
	fmt.Println(now.Unix())               //时间戳
	y, m, d := now.Local().Date()         //年月日
	hour, min, sec := now.Local().Clock() //时分秒
	fmt.Println(y, m, d)
	fmt.Println(hour, min, sec)
	fmt.Println(now.Year(), now.Month(), now.Day())
}

//格式化输出日期
func TimeFormat() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

//时间戳转日期
func TodateTime() {
	//nowTime := time.Now().Unix()
	fmt.Println(time.Unix(1513221103, 0).Format("2006-01-02 15:04:05"))
}

//日期转为时间戳
func ToTimeStamp() {
	str := "2017-12-14 11:11:43"
	//本地时间---time.local
	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Unix())
}
