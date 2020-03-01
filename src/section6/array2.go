package main

import (
	"fmt"
)

func main2() {
	// 배열 순회

	//exam 1
	arr1 := [5]int{1, 10, 100, 1000, 10000}

	//len 길이 반복

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	fmt.Println()

	arr2 := [5]int{7, 77, 777, 7777, 77777}
	for i, val := range arr2 {
		fmt.Println(i, val)
	}

	//인덱스 생략
	for _, val := range arr2 {
		fmt.Println(val)
	}

	//인덱스 생략 2
	fmt.Println()

	for i := range arr2 {
		fmt.Println(i)
	}
}
