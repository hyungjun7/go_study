package main

import "fmt"

func gcd(a int, b int) int {
	for b != 0 {
		tmp := a % b
		a, b = b, tmp
	}
	return a
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)
	fmt.Println(gcd(a, b))
	fmt.Println(lcm(a, b))
}
