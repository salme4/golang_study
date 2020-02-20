package main

import "fmt"

func main1() {
	//조건문
	//if문은 반드시 Boolean형으로 검사 -> 1, 0 (사용불가 : 자동 형 변환 불가)
	// ( ) 소괄호는 사용하지 않는다.
	var a int = 20
	b := 20

	if a > 15 {
		fmt.Println("15 이상입니다.")
	}

	if b >= 24 {
		fmt.Println("25이상이다.")
	}

	//에러 발생 1 : { } 괄호 라인 변경
	// if b >= 25
	// {
	//     fmt.Println("25이상이다.")
	// }

	//에러 발생 2 : { } 괄호 생략
	// if b >= 25
	// 	fmt.Println("25이상이다.")

	//에러 발생 3 : boolean이 아닌 식
	/*
		 if c := 1; c {
		 	fmt.Println("true")
		}
	*/

	//에러 발생 4 : 짧은 변수 scope 밖에서 사용시
	/*
		if c := 40; c >= 35 {
			fmt.Println("true")
		}
		fmt.Println("c ", c)
	*/
}
