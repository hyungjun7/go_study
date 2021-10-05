package main

import "fmt"

func main() {
	//열거형

	const (
		Jan = iota
		Feb
		Mar
		Apr
		May
	)

	//언더바 사용하면 스킵가능
	const (
		_ = iota
		A
		B
		C
		D
	)

	// 0, 1, 2, 3, 4
	fmt.Println(Jan, Feb, Mar, Apr, May)

	fmt.Println(A, B, C, D)
}
