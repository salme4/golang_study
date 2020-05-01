package main

import (
	"fmt"
)

func main5() {
	//함수 심화 1
	//가변 인자 실습(매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않음)

	//exam 1 : 가변인자 기본
	x := multi(5, 6, 7, 8, 9, 10)
	fmt.Println("ex 1 : ", x)

	y := ssum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("ex 1 : ", y)

	//exam 2 : 가변 인자 문자열
	prtWord("test1", "tt", "apple")

	//exam 3 : slice나 map을 매개변수 넘기기 ( 간편한 방법 ... 문법 활용 )
	a := []int{5, 6, 7, 8, 9, 10}

	m := multi(a...)
	n := ssum(a...)
	fmt.Println("ex 3:", m, n)
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex 2 : ", value)
	}
}

func ssum(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}

func multi(n ...int) int {
	total := 1
	for _, value := range n {
		total *= value
	}
	return total
}
