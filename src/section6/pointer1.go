package main

import (
	"fmt"
)

func main13() {
	//포인터
	//Go : 포인터 지원
	//변수의 지역성(메모리로 직접 접근 가능), 연속된 메모리 참조..., 힙, 스택
	//포인터 부분 지원(연산 X), 지역 참조성만 가져간다.
	//주소의 값은 직접 변경 불가능(잘못 된 코딩으로 버그 방지)
	//* 로 사용
	//nil로 초기화

	//exam 1
	var a *int            //방법1
	var b *int = new(int) //정석, 초기화 해서 주소값이 있음.

	fmt.Println("ex 1: ", a) //& ?
	fmt.Println("ex 1: ", b)

	i := 7

	fmt.Println("ex 1:", i, &i)

	a = &i //주소값을 전달
	b = &i

	fmt.Println("ex 1:", a, &a, *a) //i의주소, 자기 주소, (i)주소안에 값(역 참조)
	fmt.Println()
	fmt.Println("ex 1:", b, &b, *b)

	var c = &i
	d := &i

	*d = 10

	fmt.Println("ex 1:", c, &c, *c) //i의주소, 자기 주소, (i)주소안에 값(역 참조)
	fmt.Println()
	fmt.Println("ex 1:", d, &d, *d)
}
