package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = make([]int, 9, 9)
	var sum = 0
	const target, count = 100, 9
	for i := 0; i < count; i++ {
		_, _ = fmt.Scanf("%d", &arr[i])
		sum += arr[i]
	}

loop:
	for i := 0; i < count-1; i++ {
		for j := i + 1; j < count; j++ {
			if sum-(arr[i]+arr[j]) == target {
				arr[i] = -1
				arr[j] = -1
				break loop
			}
		}
	}

	sort.IntSlice(arr).Sort()

	for _, v := range arr[2:count] {
		fmt.Println(v)
	}
}
