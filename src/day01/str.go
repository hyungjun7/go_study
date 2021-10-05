package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// "", ``
	var str1 string = "c:\\go_action"
	var str2 string = `c:\go_action`
	var str3 = "hello"
	var str4 = "한글"
	var str5 = "Golang"
	var str6 = "World"

	fmt.Println(str1, str2)

	//길이 (바이트)
	fmt.Println(len(str1))

	//길이 (실제 길이)
	fmt.Println(utf8.RuneCountInString(str2))

	//아스키코드로 나옴
	fmt.Println(str3[0])
	//해결
	fmt.Printf("%c", str3[0])

	//한글 깨짐
	fmt.Printf("%c", str4[0])
	//해결
	tempStr := []rune(str4)
	fmt.Printf("%c", tempStr[0])

	//substring
	fmt.Println(str3[0:2])

	//문자열의 비교
	//아스키코드값으로 비교를 하기 때문
	fmt.Println(str5 > str6) //false

	var str7 = "asdasd"
	var str8 = "asdfsdgds"
	//문자열 연산시 + 연산보단 join 함수를 사용하는게 성능이 좋다고 함
	fmt.Println(str7 + str8)

	//기본제공하는 append
	var strSet []string
	strSet = append(strSet, str7)
	strSet = append(strSet, str8)
	fmt.Println(strSet)

	//string 패키지 join
	fmt.Println(strings.Join(strSet, ""))

}
