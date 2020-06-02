package main

import (
	"fmt"
	"reflect"
)

func main() {
	//type assertion
	//런타임 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	//인터페이스.(타입) -> 형 변환
	//interfaceValue.(type)

	//exam 1
	var a interface{} = 15
	b := a
	c := a.(int)
	//d := a.(float64)

	fmt.Println("ex 1 :", a)
	fmt.Println("ex 1 :", reflect.TypeOf(a))
	fmt.Println("ex 1 :", b)
	fmt.Println("ex 1 :", reflect.TypeOf(b))
	fmt.Println("ex 1 :", c)
	fmt.Println("ex 1 :", reflect.TypeOf(c))
	//fmt.Println("ex 1 :", d)
	fmt.Println()

	//exam 2 : 저장된 타입 검사
	if v, ok := a.(int); ok {
		fmt.Println("ex 2 :", v, ok)
	}
}
