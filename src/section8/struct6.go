package main

import (
	"fmt"
)

type Car3 struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct {
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {
	//중첩 구조체

	//exam 1
	car1 := Car3{
		"520d",
		"red",
		"bmw",
		spec{4000, 2000, 1000},
	}

	//inner struct
	fmt.Println("ex 1 :", car1.name)
	fmt.Println("ex 1 :", car1.color)
	fmt.Println("ex 1 :", car1.company)
	fmt.Printf("ex 1 : %#v\n", car1.detail)

	//inner struct field
	fmt.Println("ex 2 :", car1.detail.length)
	fmt.Println("ex 2 :", car1.detail.height)
	fmt.Println("ex 2 :", car1.detail.width)
}
