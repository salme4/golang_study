package main

import (
	"fmt"
)

func main() {
	//인터페이스 활용( empty interface )
	//empty interface : 함수 매개변수, 리턴 값, 구조체 필드 등으로 사용 가능 -> 어떤 타입으로도 변환 가능
	//모든 타입을 나타내기 위해 empty interface 사용
	//동적타입으로 생각하면 된다.( java의 Object 타입 )

	//exam 1
	var a interface{}
	printContents(a)

	a = 7.5
	printContents(a)

	a = "golang"
	printContents(a)

	a = true
	printContents(a)

	a = nil
	printContents(a)
}

func printContents(v interface{}) {
	fmt.Printf("Type : ( %T )", v)
	fmt.Println("ex 1 :", v)
}
