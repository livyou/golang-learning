package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

func RandNum() {
	rand.Seed(time.Now().UnixNano())
	randString := ""
	for i := 0; i < 6; i++ {
		//randString += strconv.Itoa(rand.Intn(9))
		randString += strconv.FormatInt(rand.Int63n(9), 10)
	}
	fmt.Println(randString)
}

//字符串转为数字---32或64位
func StringToInt() {
	strA := "123"

	intA, err := strconv.Atoi(strA)
	fmt.Println(intA, err)
	A64, err := strconv.ParseInt(strA, 10, 32)

	fmt.Println(A64, err)
}

func Md5String(t string) string {
	h := md5.New()
	h.Write([]byte(t)) // 加密字符串
	return hex.EncodeToString(h.Sum(nil))
}

func Base64Encode(t string) string {
	coder := base64.NewEncoding(base64Table)
	return coder.EncodeToString([]byte(t))
}

func Base64Decode(t string) (string, error) {
	coder := base64.NewEncoding(base64Table)
	decode, err := coder.DecodeString(t)
	return string(decode), err
}

func UrlEncode(t string) string {
	return url.QueryEscape(t)
}

func UrlDecode(t string) (string, error) {
	return url.QueryUnescape(t)
}
