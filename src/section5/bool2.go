package main

import "fmt"

func main2() {
	//exam 1
	fmt.Println("ex 1 : ", true && true)
	fmt.Println("ex 1 : ", true && false)
	fmt.Println("ex 1 : ", false && false)
	fmt.Println("ex 1 : ", true || true)
	fmt.Println("ex 1 : ", true || false)
	fmt.Println("ex 1 : ", false || false)
	fmt.Println("ex 1 : ", !true)
	fmt.Println("ex 1 : ", !false)

	//exam 2

	num1 := 15
	num2 := 37
	fmt.Println("ex 2 : ", num1 < num2)
	fmt.Println("ex 2 : ", num1 > num2)
	fmt.Println("ex 2 : ", num1 >= num2)
	fmt.Println("ex 2 : ", num1 <= num2)
	fmt.Println("ex 2 : ", num1 == num2)
	fmt.Println("ex 2 : ", num1 != num2)
}
