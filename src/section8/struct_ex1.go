package main

import (
	"fmt"
)

type Account struct {
	number   string
	balance  float64
	interest float64
}

//생성자 패턴
func newAccount(number string, balance float64, interest float64) *Account { //포인터 반환 아닌 경우 값 복사
	return &Account{number: number, balance: balance, interest: interest} //구조체 인스턴스를 생성 한 뒤 포인터를 리턴
}

func main() {
	//구조체 생성자 패턴 예제

	//exam 1
	kim := Account{number: "123-456", balance: 1000, interest: 0.015}

	var lee *Account = new(Account)
	lee.number = "234-123"
	lee.balance = 1300
	lee.interest = 0.02

	fmt.Println("ex 1 :", kim)
	fmt.Println("ex 1 :", lee)

	//exam 2 : 생성자 패턴 사용
	park := newAccount("123-333", 1700, 0.04)

	fmt.Println("ex 2 :", park)
}
