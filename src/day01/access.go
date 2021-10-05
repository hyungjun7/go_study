package main

//gofmt 가 자동으로 포맷을 하는데 _ 를 사용하면 사용하지 않는 모듈도 안없어짐
import (
	_ "day01/lib"
	chk "day01/lib2"
	"fmt"
)

func main() {
	fmt.Println(chk.CheckNum100(123))
	fmt.Println(chk.CheckNum1000(999))
}
