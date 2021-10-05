package main

import "fmt"

func main() {

	a := 10

	multi := func(x, y int) int {
		return x * y * a
	}

	fmt.Println(multi(2, 2))
}
