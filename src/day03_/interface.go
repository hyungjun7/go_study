package main

import "fmt"

func main() {
	var a interface{}
	printContent(a)
	a = 8
	printContent(a)
	a = "asdasd"
	printContent(a)
	a = false
	printContent(a)
	a = nil
	printContent(a)

	var b interface{} = 15
	c := b
	d := b.(int)
	//e := b.(float64) < 오류, 선언 당시의 타입으로만 가능

	fmt.Println(b, c, d)

	if v, ok := b.(int); ok {
		fmt.Println(v, ok)
	}
}

func chkType(arg interface{}) {
	//*.(type) 은 원래 타입을 찾아줌
	switch arg.(type) {
	case bool:
		fmt.Println(arg)
	case int8:
		fmt.Println(arg)
	case int16:
		fmt.Println(arg)
	}
}

func printContent(v interface{}) {
	fmt.Printf("%T\n", v)
}
