package main

import (
	"fmt"
	"strconv"
)

func helloGolang() {
	fmt.Println("ex 1: hello Golang")
}

func sayOne(m string) {
	fmt.Println("ex 2: ", m)
}

func sum(x int, y int) int {
	return x + y
}

func main1() {
	//함수
	//다른 언어와 차이점?? -> 반환 값(return value) 다수 가능!!
	//GO 언어만의 함수 특징
	//선언 : func 키워드로 선언
	//func 함수명(매개변수) (반환타입 or 반환값 변수명) : 반환 값 존재할 때.
	//func 함수명() (반환타입 or 반환값 변수명) : 매개변수 없음, 반환 값 존재.
	//func 함수명(매개변수) : 매개변수 존재, 반환 값 없음.
	//func 함수명() : 매개변수 없음, 반환 값 없음.

	//exam 1 : 호출
	helloGolang()

	//exam 2 : 호출 with 매개변수
	sayOne("say one!")

	//exam 3 : 호출 with 매개변수 & 반환 값
	fmt.Println("ex 3 :", sum(5, 5))
	fmt.Println("ex 3 : ", strconv.Itoa(sum(5, 5))) //숫자를 문자로 변경

}
