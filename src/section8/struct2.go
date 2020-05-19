package main

import (
	"fmt"
)

type account2 struct {
	number   string
	balance  float64 //잔액
	interest float64 //이자
}

func (a account2) calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main6() {
	//다양한 선언 방법
	//&struct, struct 차이는? : &struct 형태가 포인터를 받아오고 역참조를 하기 때문에 쪼끔 느리다.
	//Interface 메서드를 선언만 해둔 후 -> 오버라이딩 해서 메서드에 포인터 리시버를 사용할 경우 반드시 &struct형태로 넘겨야 한다.

	//포인터형 선언
	var kim *account2 = new(account2)
	kim.number = "123-456"
	kim.balance = 1000
	kim.interest = 0.015

	//
	hong := &account2{number: "123-654", balance: 15000, interest: 0.04}
	lee := new(account2)
	lee.number = "123-684"
	lee.balance = 2000
	lee.interest = 0.016

	fmt.Println("ex 1 :", kim)
	fmt.Println("ex 1 :", hong)
	fmt.Println("ex 1 :", lee)

	// 구조체 전체 출력 : %#v
	fmt.Printf("ex 2 : %#v\n", kim)
	fmt.Printf("ex 2 : %#v\n", hong)
	fmt.Printf("ex 2 : %#v\n", lee)

	fmt.Println()

	//exam 2

	fmt.Println("ex 3 :", int(kim.calculate()))
	fmt.Println("ex 3 :", int(hong.calculate()))
	fmt.Println("ex 3 :", int(lee.calculate()))
}
