package main

import "fmt"

func r77(num *int) {
	*num = 77
}

func n77(num int) {
	num = 77
}

func main() {
	var a *int
	var b = new(int)

	//<nil> 0xc00012a008
	fmt.Println(a, b)

	i := 7
	var c = &i
	fmt.Println(i, *c, &i, c)
	*c++
	fmt.Println(i, *c, &i, c)

	var d = 10
	var dd = 20
	r77(&d)
	n77(dd)

	fmt.Println(d)
	fmt.Println(dd)
}
