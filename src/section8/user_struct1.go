package main

import (
	"fmt"
)

type Car struct {
	//멤버 변수만 선언 가능
	name  string
	color string
	price int64
	tax   int64
}

//일반 메서드
func Price(c Car) int64 {
	return c.price + c.tax
}

//구조체 <-> 메서드 바인딩
func (c Car) Price() int64 {
	return c.price + c.tax
}

func main1() {
	//사용자 정의 타입
	//GO -> 객체 지향의 타입을 구조체로 정의한다.(클래스개념과, 상속 개념이 없음)
	//객체 지향에서의 클래스(속성[멤버 변수], 기능[메서드]) : 코드의 재사용성, 코드의 관리, 신뢰성 향상
	//그래도 GO는 객체 지향 언어라고 볼 수 있다.
	//GO는 전형적인 객체지향의 특징을 가지고 있진 않지만, Interface -> 다형성 지원, 구조체로 클래스형태의 코딩 가능
	//상태, 메서드 분리해서 정의한다(결합성 없음)
	//사용자 정의 타입 : 구조체, Interface, 기본타입(int, float, string), 함수까지 재정의 가능하다.
	//구조체와 메서드 연결을 통해서 타 언어의 클래스 형식처럼 사용 가능하다.(객체 지향)

	//example 1
	bmw := Car{name: "520d", price: 500000, color: "white", tax: 50000}
	benz := Car{name: "220d", price: 600000, color: "white", tax: 60000}

	fmt.Println("ex 1 :", bmw, &bmw)
	fmt.Println("ex 1 :", benz, &benz)

	fmt.Printf("ex 1 : %p", &bmw)
	fmt.Printf("ex 1 : %p", &benz)
	fmt.Println()

	fmt.Println("ex 2 :", Price(bmw))
	fmt.Println("ex 2 :", Price(benz))

	//구조체에 바인딩 된 메서드 호출
	fmt.Println("ex 3 :", bmw.Price())
	fmt.Println("ex 3 :", benz.Price())

	fmt.Println("ex 4 :", &bmw == &benz)
}
