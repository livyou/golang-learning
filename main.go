package main

import (
	"fmt"
	"golang-learning/util"
	//"sort"
	//"time"
	"runtime"
)

func main() {
	age := []int{12, 5, 21, 6, 9, 10,15}
	fmt.Printf("%T",age)
	util.SortForBubble(age)

	fmt.Println(age)
	runtime.GOMAXPROCS(2)
	util.Ttype()
	
	/*
	strs := []string{"one", "two", "three"}

	for _, s := range strs {
		go func() {
			//time.Sleep(1 * time.Second)
			fmt.Printf("%s ", s)
		}()
	}
	time.Sleep(3 * time.Second)
	stringCopy()
	util.AarraySort()

	age := []int{12, 5, 21, 6, 9, 10}
	util.SelectSort(sort.IntSlice(age))
	fmt.Println(age)
	*/
	//interfaceEqualNil()
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

func interfaceEqualNil() {
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
	fmt.Printf("%+v", userID)
}

// 拷贝的长度为两个slice中长度较小的长度值---不支持数组---有一种特殊用法，将字符串当成[]byte类型的slice
func stringCopy() {
	//有一种特殊用法，将字符串当成[]byte类型的slice
	bytes := []byte("hello world")
	copy(bytes, "ha ha")
	fmt.Println(string(bytes))

	//将第二个slice里的元素拷贝到第一个slice里，拷贝的长度为两个slice中长度较小的长度值---不支持数组
	s := []int{1, 2, 3}
	fmt.Println(s) //[1 2 3]
	copy(s, []int{4, 5, 6, 7, 8, 9})
	fmt.Println(s) //[4 5 6]
}
