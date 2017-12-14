package util

import(
	"fmt"
	"strings"
	"unicode/utf8"
	"bytes"
)

func Concat(){
	st := "hello bigticket"

	stBtye := []byte(st)
	fmt.Println(string(stBtye[1]))

	fmt.Println(strings.Join([]string{"hello"," world---1"}," "))

	for index,char := range stBtye{
		fmt.Println(index,string(char))
	}

	fmt.Println(len(st),utf8.RuneCountInString(st))

	var buf bytes.Buffer
	buf.WriteString("hello ")
	buf.WriteString("bigticket---2")
	fmt.Println(buf.String()) 

	str := "hello"
	str += " bigticket---3"

	fmt.Println(str)
}