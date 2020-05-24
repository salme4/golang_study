package main

import (
	"fmt"
)

type account struct {
	number   string
	balance  float64
	interest float64
}

func calculateD(a account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func calculateP(a *account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {

	//일반 함수 사용

	//exam 1
	kim := account{"123-111", 1000, 0.015}
	lee := account{"123-112", 2000, 0.025}

	fmt.Println("ex 1 :", kim)
	fmt.Println("ex 2 :", lee)

	fmt.Println()

	//exam 2
	calculateD(kim)
	calculateP(&lee)

	fmt.Println("ex 2 :", int(kim.balance))
	fmt.Println("ex 2 :", int(lee.balance))

}
