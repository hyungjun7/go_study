package main

import "fmt"

func main() {
	//참조 타입
	// key에는 참조값 사용 불가, value에는 가능
	var _ map[string]int = make(map[string]int)
	var map1 = make(map[string]int)

	map1["apple"] = 213
	map1["banana"] = 252

	fmt.Println(map1)

	//순회
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	//삭제
	delete(map1, "apple")

	//값 리턴시 없는 경우 해당 자료형의 기본값을 리턴함
	val, isReturn := map1["lemon"]
	//0, false
	fmt.Println(val, isReturn)
	//아래와 같은 형식으로도 사용 가능
	if v, ok := map1["lemon"]; ok {
		fmt.Println(v)
	}
}
