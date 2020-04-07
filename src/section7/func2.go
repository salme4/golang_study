package main

import (
	"fmt"
)

func main2() {
	//함수(콜백) 전달, 참조 전달(call by reference), 값 전달(call by value)

	//exam 1 : 호출 with 함수 전달
	sum2(100, add)

	//exam 2 : call by value
	a := 100
	multiValue(a)
	fmt.Println("ex 2 :", a)

	//exam 3 : call by reference
	b := 100
	multiReference(&b)
	fmt.Println("ex 3 : ", b)
}

//call by reference
func multiReference(i *int) {
	*i = *i * 10
}

//call by value
func multiValue(i int) {
	i = i * 10
}

//매개변수로 함수를 받는다.
func sum2(i int, f func(int, int)) {
	f(i, 10)
}

//매개변수 정의 시 같은 type이면 생략 가능
func add(a, b int) {
	fmt.Println("ex 1: ", a+b)
}
