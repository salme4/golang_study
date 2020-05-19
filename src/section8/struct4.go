package main

import (
	"fmt"
)

func main8() {
	//구조체 익명 선언 및 활용
	//type 구조체명 struct

	car1 := struct{ name, color string }{"520d", "red"}

	fmt.Println("ex 1 :", car1)
	fmt.Printf("ex 1 : %#v\n", car1)

	//exam 2 반복문
	cars := []struct{ name, color string }{{"520d", "red"}, {"530d", "white"}, {"540d", "blue"}}

	for _, c := range cars {
		fmt.Printf("ex 2 : (%s, %s) ---- (%#v)\n", c.name, c.color, c)
	}
}
