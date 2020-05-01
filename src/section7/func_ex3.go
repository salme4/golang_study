package main

import (
	"fmt"
)

func main7() {
	//함수 심화 3
	//재귀 함수

	//ex 1 : 팩토리얼
	x := fact(7)
	fmt.Println("ex : 1", x)

	//ex 2 : print 재귀함수
	prtHello(10)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("ex 2: (", n, ")", "hello!")
	prtHello(n - 1)
}
