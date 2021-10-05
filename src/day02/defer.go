package main

import "fmt"

func f1() {
	fmt.Println("asdasdasd")
	defer f2()
}

func f2() {
	fmt.Println("exit")
}

func f3() {
	defer func() {
		fmt.Println("asdasd")
	}()
}

//FILO
func stack() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func s(str string) string {
	fmt.Println("s")
	return str
}

func f(str string) {
	fmt.Println("f")
}

func a() {
	//defer 는 바로 뒤에 있는 함수만 지연실행시킴 나머지는 포함안됨
	defer f(s("asdasd"))
	fmt.Println("111")
}

func main() {
	//defer
	//defer 호출한 함수가 종료되기 직전에 호출함, finally
	f1()
	f3()
	stack()
	a()
}
