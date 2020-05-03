package main

import (
	"fmt"
)

func main11() {
	//ex 1
	stack()
}

func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("ex 1 :", i) //FILO
	}
}
