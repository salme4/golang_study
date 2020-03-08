package main

import (
	"fmt"
)

func main5() {
	//슬라이스 참조 타입 증명

	//exam 1
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1 //값 복사
	arr2[0] = 7

	fmt.Println("ex 1 : ", arr1)
	fmt.Println("ex 1 : ", arr2)
	fmt.Println()

	//exam 2
	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1
	slice2[0] = 7

	fmt.Println("ex 2 : ", slice1)
	fmt.Println("ex 2 : ", slice2)
	fmt.Println()

	//exam 3 : exception 길이만큼 초기화
	sliec3 := make([]int, 50, 100)
	fmt.Println("ex 3 : ", sliec3[4])
	//fmt.Println("ex 3 : ", sliec3[50]) // 길이 index orver
	fmt.Println()

	//exam 4 : for
	for i, v := range slice1 {
		fmt.Println("ex 4: ", i, v)
	}
}
