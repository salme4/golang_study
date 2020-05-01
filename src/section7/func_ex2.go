package main

import (
	"fmt"
)

func main6() {
	//함수를 변수에 할당하여 사용.

	//ex 1 : slice에 할당 (함수타입의 슬라이스 생성 및 초기화)
	f := []func(int, int) int{multiply2, sum3}

	a := f[0](2, 3)
	b := f[1](2, 3)

	fmt.Println("ex 1 :", a, f[0](2, 3))
	fmt.Println("ex 1 :", b, f[1](2, 3))

	//ex 2 : 변수에 할당
	var f1 func(int, int) int = multiply2
	f2 := sum3

	fmt.Println("ex 2 :", f1(2, 3))
	fmt.Println("ex 2 :", f2(2, 3))

	//ex 3 : map에 할당
	m := map[string]func(int, int) int{
		"mul_func": multiply2,
		"sum_func": sum3,
	}

	fmt.Println("ex 3 :", m["mul_func"](2, 3))
	fmt.Println("ex 3 :", m["sum_func"](2, 3))
}

func multiply2(x, y int) (r int) {
	r = x * y
	return r
}

func sum3(x, y int) (r int) {
	r = x + y
	return r
}
