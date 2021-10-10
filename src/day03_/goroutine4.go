package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//고루틴과 클로저
	runtime.GOMAXPROCS(1)

	s := "closure"

	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, "====", time.Now())
		}(i) //고루틴 클로저는 가장 나중에 실행됨
	}
	time.Sleep(4 * time.Second)
}
