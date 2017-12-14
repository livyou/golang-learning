package main

import (
	"fmt"
	"golang-learning/util"
)

func main() {

	interfaceEqualNil()
	//util.Concat()
	/*
	fmt.Println("---here---")
	testRand()
	fmt.Println("---here---")
	testTime()
	*/
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

func interfaceEqualNil(){
	var userID interface{}
    userID = 123

    var id int32
    id = 123
    
    //这里不能赋值，因为类型不一样
    //id = user_id

    //但是这里可以判断，为什么不同的类型可以判断相等？？？
    if userID == id {
        fmt.Println("相等", userID)
    } else {
        fmt.Println("不相等", userID)
	}
	fmt.Printf("%+v",userID)
}
