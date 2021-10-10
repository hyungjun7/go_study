package main

import (
	"fmt"
	"time"
)

func exec() {
	fmt.Println("exec start")
	time.Sleep(1 * time.Second)
	fmt.Println("exec end")
}

func exec2() {
	fmt.Println("exec2 start")
	time.Sleep(1 * time.Second)
	fmt.Println("exec2 end")
}

func exec3() {
	fmt.Println("exec3 start")
	time.Sleep(1 * time.Second)
	fmt.Println("exec3 end")
}

func main() {
	exec()

	fmt.Println("main routine start", time.Now())
	//고루틴 실행
	go exec2()
	go exec3()
	time.Sleep(4 * time.Second)
	fmt.Println("main routine end", time.Now())
}
