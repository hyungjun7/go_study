package main

import "fmt"

type Car struct {
	name  string
	color string
	price int64
}

// Price 일반 메소드
//func Price(c Car) int64 {
//	return c.price
//}

// Price 구조체에 메소드 바인딩
func (a Car) Price() int64 {
	return a.price
}

func tt(n int) {
	fmt.Println(n)
}

func main() {
	//구조체와 메소드를 연결해서 클래스 흉내낼 수 있음
	kia := Car{
		name:  "k3",
		price: 1200203,
		color: "red",
	}

	fmt.Println(kia.Price())

	type cnt int

	a := cnt(15)
	b := 15
	//tt(a) 에러남, 형변환 안해줌
	tt(b)
	tt(int(a))
}
