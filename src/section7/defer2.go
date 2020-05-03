package main

import (
	"fmt"
)

func main10() {
	sayHello("Golang")
	fmt.Println("end")
}

func sayHello(str string) {
	defer func() {
		fmt.Println("str : ", str)
	}()

	func() {
		fmt.Println("hi ")
	}()
}
