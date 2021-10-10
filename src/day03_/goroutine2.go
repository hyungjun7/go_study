package main

import (
	"fmt"
	"time"
)

func exec6(name string) {
	fmt.Println(name, "start", time.Now())
	for i := 0; i < 1000; i++ {
		fmt.Println(name, i)
	}
	fmt.Println(name, "end", time.Now())
}

func main() {
	exec6("t1")

	fmt.Println("main routine start", time.Now())
	go exec6("t2")
	go exec6("t3")
	time.Sleep(4 * time.Second)
	fmt.Println("main routine end", time.Now())
}
