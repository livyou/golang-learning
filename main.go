package main

import (
	"fmt"
	"learn/util"
)

func main() {
	fmt.Println("---here---")
	testRand()
	fmt.Println("---here---")
	testTime()
}

func testRand() {
	util.RandNum()
	util.StringToInt()
	url := util.UrlEncode("http://www.163.com")
	fmt.Println(url)
	urlDecode, _ := util.UrlDecode(url)
	fmt.Println(urlDecode)

	str := util.Md5String("123456")
	fmt.Println(str)

	base64Str := util.Base64Encode("123456")
	fmt.Println(base64Str)
	base64DecodeStr, _ := util.Base64Decode(base64Str)
	fmt.Println(base64DecodeStr)
}

func testTime() {
	util.TimeStamp()
	util.TodateTime()
	util.TimeFormat()
}
