package main

import "fmt"

func main()  {
	const (
		_ = iota
		A
		_
		C
	)
	// _ 는 skip하기 위해 사용 할 수 있다.
	fmt.Println(A, C)

	const (
		_ = iota + 0.75 * 2
		DEFAULT
		SILVER
		_
		PLATINUM
	)

	fmt.Println(DEFAULT, SILVER, PLATINUM)
}