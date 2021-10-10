package main

import "fmt"

func main() {
	var a, b, v int
	_, _ = fmt.Scanln(&a, &b, &v)
	fmt.Println((v-b-1)/(a-b) + 1)
}
