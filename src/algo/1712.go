package main

import "fmt"

func main() {
	var a, b, c int
	_, _ = fmt.Scanln(&a, &b, &c)
	if b > c || b == c {
		fmt.Println(-1)
		return
	}
	fmt.Println(a/(c-b) + 1)
}
