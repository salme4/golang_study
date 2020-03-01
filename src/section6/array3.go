package main

import (
	"fmt"
)

func main3()  {
	//배열 값 복사 증명
	
	//exam 1
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	arr2 := arr1

	//& : 주소의 값(?)을 출력
	fmt.Println("ex 1 :", arr1, &arr1)
	fmt.Println("ex 1 :", arr2, &arr2)

	fmt.Printf("ex 1 : %p %v\n", &arr1, arr1)
	fmt.Printf("ex 1 : %p %v\n", &arr2, arr2)
}
