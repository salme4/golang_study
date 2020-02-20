//열거형 2
package main

import "fmt"

func main7() {
	const (
		A = iota * 10
		B
		C
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
