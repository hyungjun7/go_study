package main

import "fmt"

func main() {
	var num, num2, num3 int = 1, 2, 3
	var str string = "123124"
	var float float32 = 3.14
	var str2 = "a21523"

	const myname = "hyungjun"

	var (
		name      = "mac"
		version   = 2020
		isRunning = false
	)

	//짧은 선언
	//선언 후에 재할당하면 오류남
	short := 3

	fmt.Println(num, num2, num3)

	fmt.Println(myname)

	fmt.Println(str, float, str2)

	fmt.Println(name, version, isRunning)

	fmt.Println(short)

	if i := 10; i < 12 {
		fmt.Print("10보다 작네")
	} else {
		fmt.Print("12보다 큼")
	}
}
