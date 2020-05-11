package main

import (
	"fmt"
)

type cnt int

func main2() {
	//사용자 정의 타입(기본 자료형 타입)

	//exam 1 : 선언
	a := cnt(15)
	fmt.Println("ex 1:", a)

	//exam 2
	var b cnt = 15
	fmt.Println("ex 1:", b)

	//exam 3
	//testConvertT(b) //사용자 정의 타입과 <-> 기본 타입 : 매개변수 전달 시에 변환해야 사용 가능
	testConvertT(int(b)) //형 변환
	testConvertD(b)
}

func testConvertT(i int) {
	fmt.Println("ex 3: (Default Type) :", i)
}

func testConvertD(i cnt) {
	fmt.Println("ex 3: (Custom Type) :", i)
}
