package main

import (
	"fmt"
)

func main7()  {
	//exam 1 
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println("ex 1 : ", n1+n2)
	fmt.Println("ex 1 : ", n1-n2)
	fmt.Println("ex 1 : ", n1*n2)
	fmt.Println("ex 1 : ", n1/n2)
	fmt.Println("ex 1 : ", n1%n2)
	fmt.Println("ex 1 : ", n1<<2)
	fmt.Println("ex 1 : ", n1>>2)
	fmt.Println("ex 1 : ", ^n1)

	var n3 int =12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	fmt.Println("ex 2 : ", float32(n3) + n4)
	fmt.Println("ex 2 : ", n3 + int(n4))

	fmt.Println("ex 2 : ", n5 + uint16(n6)) //type orverflow 때문에 type 의 Max 값으로 계산된다.
}