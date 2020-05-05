package main

import (
	"fmt"
)

func main() {

	//ex 1
	cnt := increseCnt()

	fmt.Println("ex 1 :", cnt)
	fmt.Println("ex 1 :", cnt)
	fmt.Println("ex 1 :", cnt)
	fmt.Println("ex 1 :", cnt)
	fmt.Println("ex 1 :", cnt)
}

func increseCnt() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
