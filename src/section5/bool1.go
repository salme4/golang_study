package main

import "fmt"

func main1() {
	//Bool : 참과 거짓
	//조건부 논리 연사자랑 주로 사용 : !, ||, &&
	//암묵적 형 변환 X : 0, nil -> false가 아니다.

	//exam 1
	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println("ex 1 : ", b1)
	fmt.Println("ex 2 : ", b2)
	fmt.Println("ex 3 : ", b3)

	fmt.Println("ex 4 : ", b3 == b3)
	fmt.Println("ex 5 : ", b1 && b3)
	fmt.Println("ex 6 : ", b1 || b2)
	fmt.Println("ex 7 : ", !b1)

	//에러 발생
	/*
		b4 := 1
		if b4 {
			fmt.Println("ex 8 : true")
		}
	*/
}
