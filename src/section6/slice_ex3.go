package main

import (
	"fmt"
)

func main8() {
	//슬라이스 복사
	//copy(복사대상, 원본)
	//make() 로 공간 할당 후 복사해야 한다.
	//복사 된 슬라이스 값 변경해도 원본에는 영향이 없음.

	//exam 1
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1) //5개만 복사 됨.
	copy(slice3, slice1) //복사 안됨.

	fmt.Println("ex 1 :", slice2)
	fmt.Println("ex 1 :", slice3)
	fmt.Println()

	//exam 2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)
	copy(b, a)

	b[0] = 7
	b[4] = 10

	fmt.Println("ex 2 :", a)
	fmt.Println("ex 2 :", b)
	fmt.Println()

	//exam 3
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3] //추출은 값 복사가 아니라 참조 복사이다. -> 원본 값 변경됨!!!

	d[1] = 7

	fmt.Println("ex 3 :", c)
	fmt.Println("ex 3 :", d)
	fmt.Println()

	//exam 4
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] //마지막은 용량을 7로 세팅한다.

	fmt.Println("ex 4 : ", len(f), cap(f))
	fmt.Println("ex 4 : ", f)

}
