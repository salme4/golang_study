package main

import "fmt"

func main8() {
	//반복문
	//go에서 유일하게 제공하는 반복문
	//다양한 사용법 숙지

	//exam 1
	for i := 0; i < 5; i++ {
		fmt.Println("ex 1 ", i)
	}

	//에러 발생 1 : { } 괄호 줄 바꿈
	/*
		for i := 0; i < 5; i++
		{
			fmt.Println("ex 1 ", i)
		}
	*/

	//에러 발생 2 : { } 괄호 생략
	/*
		for i := 0; i < 5; i ++
		fmt.Println("ex 1 ", i)
	*/

	//exam 2 : 무한 루프
	/*
		for {
			fmt.Println("ex 2 : Hello golang!")
			fmt.Println("ex 2 : Infinite Loop")
		}
	*/

	//exam 3 : Range 용법
	location := []string{"Seoul", "Busan", "Incheon"}
	for _, name := range location {
		fmt.Println("ex 3 : ", name)
	}
}
