package main

import (
	"fmt"
)

func main11() {
	//문자열 연산
	//추출, 비교, 조합

	//exam1
	var str1 string = "Golang"
	var str2 string = "world"

	fmt.Println("ex 1 : ", str1[0:2], str1[0])
	fmt.Println("ex 1 : ", str2[3:], str2[0])
	fmt.Println("ex 1 : ", str2[:4])
	fmt.Println("ex 1 : ", str1[1:3])
}
