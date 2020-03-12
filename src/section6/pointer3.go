package main

import (
	"fmt"
)

func rptc(n *int) {
	*n = 77
}

func vptc(n int) {
	n = 77
}

func main() {
	//포인터의 값 전달
	//함수, 메서드 호출 시에 매개변수 값을 복사 전달 -> 함수, 메서드 내에서는 원칙적으로 원본 값 변경 불가능.
	//원본 값 변경을 위해 포인터로 전달한다.
	//특히 크기가 큰 배열인 경우 값 복사 시 부담 -> 포인터 전달로 해결(슬라이스, 맵은 참조 전달)

	//exam 1

	var a int = 10
	var b int = 10

	fmt.Println("ex 1:", a)
	fmt.Println("ex 1:", b)
	fmt.Println()

	rptc(&a)
	vptc(b)

	fmt.Println("ex 1:", a)
	fmt.Println("ex 1:", b)
	fmt.Println()
}
