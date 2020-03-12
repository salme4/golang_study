package main

import (
	"fmt"
)

func main14() {
	//exam 1

	i := 7
	p := &i

	fmt.Println("ex 1:", i, *p, &i, p)

	*p++ //역참조 값 증가

	fmt.Println("ex 1:", i, *p, &i, p)

	*p = 77777 //역참조 값 변경

	fmt.Println("ex 1:", i, *p, &i, p)

	i = 77 //원본 값 변경

	fmt.Println("ex 1:", i, *p, &i, p)
}
