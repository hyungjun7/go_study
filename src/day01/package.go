package main

import (
	"day01/lib"
	"fmt"
)

func main() {
	//같은 패키지 내부에 있는 소스파일들은 디렉토리명을 패키지명으로 사용함
	//접근 제어자는 소문자는 private, 대문자는 public
	fmt.Println(lib.CheckNum(8))
}
