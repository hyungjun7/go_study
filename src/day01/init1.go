package main

import "fmt"

//init 은 메인보다 먼저 실행된다.
//여러개 있어도 상관없음, 위에서 부터 아래로 차례대로 실행
//임포트한 패키지에 init이 있다면 그것도 실행된다
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
