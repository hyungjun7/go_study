package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func exec4(name int) {
	r := rand.Intn(100)
	fmt.Println(name, "start", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>", r, i)
	}
	fmt.Println(name, "end", time.Now())
}

func main() {
	//멀티코어 활용
	//사용할 코어 설정 //실행할 수 있는 코어 리턴
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("main routine start", time.Now())

	for i := 0; i < 100; i++ {
		go exec4(i) //고루틴  100개
	}

	time.Sleep(5 * time.Second)

	fmt.Println("core", runtime.GOMAXPROCS(0))

	fmt.Println("main routine end", time.Now())
}
