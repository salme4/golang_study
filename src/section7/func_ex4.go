package main

import (
	"fmt"
)

func main()  {
	//함수 심화 4
	//익명 함수
	//즉시 실행( 선언과 동시에)
	
	//ex 1
	func() {
		fmt.Println("ex : 1 anonymous Function!")
	}()

	//ex 2
	msg := "Hello golang@"
	func(m string) {
		fmt.Println("ex 2:", m)
	}(msg)

	//ex 3
	func (x, y int)  {
		fmt.Println("ex 3:", x + y)
	}(10, 20)

	//ex 4
	r := func(x, y int) int {
		return x * y
	}(10, 100)
	fmt.Println("ex 4 :", r)
}