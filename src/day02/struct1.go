package main

import (
	"fmt"
	"reflect"
)

type Account struct {
	id       string
	balance  float64
	interest float64
}

func (a Account) Calc() float64 {
	return a.balance + (a.balance * a.interest)
}

//필드에 태그 설정 가능
type Car2 struct {
	name    string "이름"
	color   string "색"
	company string "회사"
}

func main() {
	kim := Account{
		id:       "1235235",
		balance:  100000,
		interest: 0.3,
	}

	lee := Account{
		id:       "1232532",
		interest: 0.4,
	}

	//초기화하지 않은 값은 기본값으로 초기화된다,
	fmt.Println(kim, lee)

	fmt.Println(int(kim.Calc()))

	//아래와 같이 선언가능
	//위 보다 속도가 조금 느리지만 특정 상황(인터페이스)에서는 이렇게 사용해야 함
	var park *Account = new(Account)
	park.id = "123124123"
	park.balance = 100010
	park.interest = 0.4

	hong := &Account{
		id:       "1233546",
		balance:  10000,
		interest: 0.09,
	}
	fmt.Println(hong)

	//구조체 익명선언
	car2 := struct {
		name  string
		color string
	}{
		"casper",
		"blue",
	}
	fmt.Printf("%#v\n", car2)

	tag := reflect.TypeOf(Car2{})

	for i := 0; i < tag.NumField(); i++ {
		fmt.Println(tag.Field(i).Tag)
	}
}
