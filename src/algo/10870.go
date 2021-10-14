package main

import "fmt"

func getFibonacci(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return getFibonacci(num-1) + getFibonacci(num-2)
	}
}

func main() {
	var num int = 0
	_, _ = fmt.Scanf("%d", &num)
	if num == 0 {
		fmt.Println(0)
		return
	}
	tmp := getFibonacci(num)
	fmt.Println(tmp)
}
