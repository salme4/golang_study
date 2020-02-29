package main

import (
	"fmt"
)

func main12() {

	//문자열 비교
	str1 := "Golang"
	str2 := "World"

	fmt.Println("ex 1 :", str1 == str2)
	fmt.Println("ex 2 :", str1 != str2)
	fmt.Println("ex 1 :", str1 > str2)
	fmt.Println("ex 1 :", str1 < str2) //Go 문자열 비교 시 ASCII 코드값으로 사전 식 비교

}
