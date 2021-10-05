package main

import "fmt"

func main() {
	a := -7

	//break 안붙여도 된다.
	//만약 조건이 참이더라도 다음 케이스 문으로 들어간다고 하면 fallthrough 키워드를 붙여주면 된다.
	switch {
	case a < 0:
		fmt.Println("a < 0")
		fallthrough
	case a == 0:
		fmt.Println("a == 0")
	case a > 0:
		fmt.Println("a > 0")
	}
}
