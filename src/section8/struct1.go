package main

import (
	"fmt"
)

type account struct {
	number   string
	balance  float64 //잔액
	interest float64 //이자
}

//Account Receiver Method
func (a account) calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	//구조체
	//GO 의 필드들의 집합체 또는 컨테이너
	//필드는 갖지만 메서드는 갖지 않는다.
	//객체지향 방식을 지원 -> 리시버를 통해 메서드랑 연결
	//상속, 객체, 클래스 개념이 없다.
	//구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	//exam 1 선언방법
	kim := account{number: "123-456", balance: 10000000, interest: 0.015}
	lee := account{number: "123-543", balance: 1200}
	park := account{number: "123-584", interest: 0.025}
	cho := account{"123-985", 1500, 0.03}

	fmt.Println("ex 1 :", kim)
	fmt.Println("ex 1 :", lee)
	fmt.Println("ex 1 :", park)
	fmt.Println("ex 1 :", cho)

	fmt.Println()

	//exam 2 : receiver call
	fmt.Println("ex 2 :", int(kim.calculate()))
	fmt.Println("ex 2 :", int(lee.calculate()))
	fmt.Println("ex 2 :", int(park.calculate()))
	fmt.Println("ex 2 :", int(cho.calculate()))

}
