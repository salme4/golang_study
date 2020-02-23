package main

import (
	"fmt"
)

//init 여러개는 별로 권고하지 않는다.
func init() {
	fmt.Println("Init1 Method Start.")
}

func init() {
	fmt.Println("Init2 Method Start.")
}

func init() {
	fmt.Println("Init3 Method Start.")
}

func init() {
	fmt.Println("Init4 Method Start.")
}

func main6() {
	fmt.Println("main start")
}
