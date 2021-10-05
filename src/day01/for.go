package main

import "fmt"

func main() {

	//while
	for {
		fmt.Println("asdasd")
	}

	//레이블 밑에는 for 만 올 수 있음
Loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				//아래와 같은 방법으로 한 번에 2중 포문을 탈출 가능 continue 도 마찬가지임
				break Loop
			}
		}
	}
}
