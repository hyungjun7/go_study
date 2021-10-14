package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scanln(&a, &b)
	var arr = make([]int, 0, a)
	for i := 1; i < a; i++ {
		if a%i == 0 {
			arr = append(arr, i)
		}
	}
	arr = append(arr, a)
	if b > len(arr) {
		fmt.Println(0)
		return
	}
	fmt.Println(arr[b-1])
}
