package main

import (
	"fmt"
	"strconv"
)

func sum(x int, y int) int {
	return x + y
}

//cb
func sum2(x int, f func(int, int) int) int {
	return f(1, 2) + x
}

func referrer(x *int) int {
	return *x * *x
}

//값 2개 리턴
func m(x, y int) (int, int) {
	return x * 10, y * 10
}

func h(x, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y * 10
	return r1, r2
}

//가변인자
func multi(n ...int) int {
	tot := 1
	for _, v := range n {
		tot *= v
	}
	return tot
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(sum(1, 2))
	//toString
	fmt.Println(strconv.Itoa(sum(1, 2)))

	fmt.Println(sum2(100, sum))

	c := 7
	fmt.Println(referrer(&c))

	y, z := m(c, c)

	fmt.Println(y, z)

	p, pp := h(c, c)
	fmt.Println(p, pp)

	fmt.Println(multi(1, 2, 3, 4, 5))

	ap := []int{1, 2, 3, 4, 5}
	fmt.Println(multi(ap...))

	ff := []func(int, int) (int, int){m, h}
	fmt.Println(ff[0](1, 2))

	var f1 = m
	fmt.Println(f1(1, 2))

	asd := map[string]func(int, int) (int, int){
		"m": m,
		"h": h,
	}
	fmt.Println(asd["m"](1, 2))

	x := fact(7)
	fmt.Println(x)

	func(b int) {
		fmt.Println(b)
	}(10)
}
