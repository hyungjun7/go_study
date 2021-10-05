package main

import "fmt"

func main() {
	var a = 15
	b := 12

	//중괄호를 C언어 처럼 하면 에러남
	//if문 괄호를 생략해도 에러
	if a < b {
		fmt.Println("ㅁㄴㅇㅁㄴㅇ")
	}
}
