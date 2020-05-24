package main

import (
	"fmt"
)

type account struct {
	number   string
	balance  float64
	interest float64
}

func (a account) calculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *account) calculateP(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {

	//구조체 인스턴스 값 변경 시 -> 포인터 전달, 보통의 경우 -> 값 전달
	//receiver 활용

	//exam 1
	kim := account{"123-111", 1000, 0.015}
	lee := account{"123-112", 2000, 0.025}

	fmt.Println("ex 1 :", kim)
	fmt.Println("ex 2 :", lee)

	fmt.Println()

	lee.calculateD(100)
	kim.calculateP(100)

	fmt.Println("ex 2 :", int(lee.balance))
	fmt.Println("ex 2 :", int(kim.balance))
}
