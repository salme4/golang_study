package main

import (
	"fmt"
)

func main10() {
	//문자열 표현
	//문자열 : 기본 UTF-8 인코딩 (유니코드 문자 집합)

	//exam 1
	var str1 string = "Golang"
	var str2 string = "World"
	var str3 string = "고프로그래밍"

	fmt.Println("ex 1 : ", str1[0], str1[1], str1[2], str1[3]) //ascii code 값 출력
	fmt.Println("ex 1 : ", str2[0], str2[1], str2[2], str2[3])
	fmt.Println("ex 1 : ", str3[0], str3[1], str3[2], str3[3])

	//exam 2
	fmt.Println()
	fmt.Printf("ex 2 : %c, %c, %c, %c \n", str1[0], str1[1], str1[2], str1[3])
	fmt.Printf("ex 2 : %c, %c, %c, %c \n", str2[0], str2[1], str2[2], str2[3])
	fmt.Printf("ex 2 : %c, %c, %c, %c \n", str3[0], str3[1], str3[2], str3[3]) //한글 깨짐

	conStr := []rune(str3) //format 형으로 한글 출력하기
	fmt.Printf("ex 2 : %c, %c, %c, %c \n", conStr[0], conStr[1], conStr[2], conStr[3])

	//exam 3
	for i, char := range str1 {
		fmt.Printf("ex 3: %c(%d)\t", char, i)
	}
	fmt.Println()

	for i, char := range str2 {
		fmt.Printf("ex 4: %c(%d)\t", char, i)
	}
}
