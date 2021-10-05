package main

import (
	"fmt"
	"sort"
)

func main() {
	//가변길이, 참조형

	//사이즈를 할당하지 않으면 slice 임
	var _ []int
	_ = []int{}
	_ = []int{1, 2, 3, 4, 5}
	_ = [][]int{}

	//길이가 없기 때문에
	//slc1[0] = 2 < 에러

	//make(t, size, cap)
	//사이즈는 인덱스로 접근가능, 용량은 메모리에 공간 확보, 용량 생략시 사이즈만큼
	var _ []int = make([]int, 5, 10)

	//nil
	var _ []int

	//참조 확인
	arr1 := []int{1, 2, 3, 4, 5}
	var arr2 = arr1
	arr2[2] = 1
	fmt.Println(arr1, arr2)

	//순회
	for v := range arr1 {
		fmt.Println(v)
	}

	//push
	arr2 = append(arr2, 6, 7, 8)

	//join, 전개연산자를 써줘야 함
	var arr3 = []int{10, 11, 12}
	arr3 = append(arr2, arr3...)
	arr3 = append(arr2, arr3[0:1]...)

	//sort
	var arr4 = []int{10, 11, 12, 1, 6, 8, 19}
	//정렬이 되어 있는지 확인
	fmt.Println(sort.IntsAreSorted(arr4))
	sort.Ints(arr4)
	fmt.Println(arr4)

	//copy
	//make 로 공간 확보 후 복사해야함, 넘치는 경우 넘치는 건 자동 제외
	//원본 배열을 복사해서 새로운 배열을 만들어냄
	s := []int{1, 2, 3, 4}
	target := make([]int, 5)
	copy(target, s)

	_ = s[0:2]   //이거는 참조상태임
	_ = s[0:2:7] //용량 복사 가능
}
