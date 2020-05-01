package main

import (
	"fmt"
)

func main4() {
	//변수지정 리턴

	//exam 1
	a, b := namingFunc(10, 20)
	fmt.Println("ex 1 :", a, b)
}

//반환값에 이름을 지정하여 가독성을 높임.
func namingFunc(x, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y * 20
	return r1, r2
}
