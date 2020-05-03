package main

import (
	"fmt"
)

func main() {
	//ex 1
	a()
}

func a() {
	defer end(start("b")) //defer로 지정한 end 함수만 지연된다. 중첩 함수 주의!
	fmt.Println("in a")
}

func start(str string) string {
	fmt.Println("start ", str)
	return str
}

func end(str string) {
	fmt.Println("end : ", str)
}
