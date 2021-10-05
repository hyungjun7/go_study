package main

import "fmt"

func main() {
	//배열과 슬라이스
	//길이 고정 / 길이 가변
	//val / ref
	//복사값 / 참조값
	//비교 연산자 가능 / 불가

	//cap() 배열, 슬라이스 용량
	//len() 배열 슬라이스 갯수

	var _ [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	//배열 크기 자동 맞춤
	_ = [...]int{1, 2, 3, 4, 5}
	_ = [5][5]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}

	//배열 순회
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	for idx, val := range arr2 {
		fmt.Println(idx, val)
	}

	for v := range arr2 {
		fmt.Println(v)
	}

	//배열 복사 확인
	tempArr := arr2
	//0xc0000b4030, 0xc0000b4000
	//값을 복사하기 때문에 메모리에 배열이 하나 더 생김
	fmt.Printf("%p, %p", &tempArr, &arr2)
}
