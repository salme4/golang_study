package main

import (
	"fmt"
)

type employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e employee) calculate() float64 {
	return e.salary + e.bonus
}

type executives struct {
	employee     //is a 관계 -> 임원도 직원이다
	specialBonus float64
}

func main() {
	//구조체 임베디드 패턴을 사용함으로써,
	//다른 관점으로 메서드를 재 사용하는 장점 제공
	//상속을 허용하지 않는 Go 언어에서 메서드 상속 활용을 위한 패턴
	//상속

	//exam 1 : 직원
	ep1 := employee{"kim", 200, 10}
	ep2 := employee{"park", 150, 20}

	//임원
	ex := executives{
		employee{"lee", 500, 100},
		100,
	}

	fmt.Println("ex 1 :", int(ep1.calculate()))
	fmt.Println("ex 1 :", int(ep2.calculate()))

	//employee 구조체 통해서 메서드 호출
	fmt.Println("ex 1 :", int(ex.calculate()+ex.specialBonus))
}
