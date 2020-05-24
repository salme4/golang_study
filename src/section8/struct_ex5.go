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

func (e executives) calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type executives struct {
	employee
	specialBonus float64
}

func main() {
	//구조체 임베디드 메서드 오버라이딩 패턴
	//오버라이딩

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
	fmt.Println("ex 1 :", int(ex.calculate()))
	fmt.Println("ex 1 :", int(ex.employee.calculate()+ex.specialBonus))

}
