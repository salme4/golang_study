package main

import (
	"fmt"
)

func main4() {
	//길이가 가변이다. -> 동적으로 크기가 늘어난다. 참조타입이다.
	//슬라이스(길이 & 용량) 크기가 동적으로 할당 가능.
	//슬라이스를 대부분 사용한다.

	//2가지 선언 방법 : 배열처럼, make 함수 사용 : make(자료형, 길이, 용량(option))

	//exam 1
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}

	//slice2[0] = 1 <- 용량이 초기화 되지 않음
	slice3[4] = 10 //값 수정 가능

	fmt.Printf("ex 1 : %-5T %d %d %v \n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("ex 1 : %-5T %d %d %v \n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("ex 1 : %-5T %d %d %v \n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("ex 1 : %-5T %d %d %v \n", slice4, len(slice4), cap(slice4), slice4)

	//exam 2 : 타입, 길이, 용량을 확보
	var slice5 []int = make([]int, 5, 10)
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)

	slice6[2] = 7

	fmt.Printf("ex 2 : %-5T %d %d %v \n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("ex 2 : %-5T %d %d %v \n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("ex 2 : %-5T %d %d %v \n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("ex 2 : %-5T %d %d %v \n", slice8, len(slice8), cap(slice8), slice8)

	//exam 3
	var slice9 []int //nil 슬라이스 (길이와 용량이 0)
	if slice9 == nil {
		fmt.Println("this is nil slice")
	}
}
