package main

import (
	"fmt"
)

func main3() {
	//다중 값 반환

	//exam 1
	a, b := multiply(10, 5)
	//c := multiply(10, 5)
	c, _ := multiply(10, 5)
	_, d := multiply(10, 5)

	fmt.Println("ex 1:", a, b, c, d)

	//exam 2
	x1, x2, x3, x4, x5 := arrayMultiply(1, 2, 3, 4, 5)

	fmt.Println("ex 2:", x1, x2, x3, x4, x5)
}

func multiply(x, y int) (int, int) {
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, d * 4, e * 5
}
