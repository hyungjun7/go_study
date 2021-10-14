package main

import "fmt"

func main() {
	var in, out, max, now int
	for i := 0; i < 10; i++ {
		_, _ = fmt.Scanf("%d %d", &out, &in)
		now += in - out
		if max < now {
			max = now
		}
	}
	fmt.Println(max)
}
