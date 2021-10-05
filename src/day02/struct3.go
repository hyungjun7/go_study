package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) goOut() {
	fmt.Println("편의점가야지")
}

//상속, 구조체 임베디드 패턴
type Student struct {
	Human
	school string
}

//override
func (s Student) goOut() {
	fmt.Println("학교가야지")
}

func main() {

}
