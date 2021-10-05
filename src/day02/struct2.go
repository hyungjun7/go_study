package main

import "fmt"

type Car3 struct {
	name    string "이름"
	color   string "색"
	company string "회사"
	detail  spec   "디테일"
}

type spec struct {
	length int
	height int
	width  int
}

func newCar(name, color, company string, detail spec) *Car3 {
	return &Car3{name, color, company, detail}
}

//중첩 구조체
func main() {
	car := Car3{
		name:    "gsdsa",
		color:   "blue",
		company: "kia",
		detail: spec{
			length: 12,
			height: 13,
			width:  1,
		},
	}

	fmt.Printf("%#v\n", car)

	tico := newCar("asd", "asd", "asd", spec{
		length: 1,
		height: 2,
		width:  3,
	})

	fmt.Println(tico)
}
