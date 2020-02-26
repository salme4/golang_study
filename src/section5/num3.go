package main

import "fmt"

func main5() {
	//부동소수점 ( 실수 )
	//float32(7자리), float64(15자리)

	//예제 1
	var num1 float32 = 0.14
	var num2 float32 = .75647
	var num3 float32 = 442.0378373
	var num4 float32 = 10.0

	//지수 표기법
	var num5 float32 = 14e6
	var num6 float32 = .156875E+3
	var num7 float32 = 5.32521e-10

	fmt.Println("ex 1 : ", num1)
	fmt.Println("ex 1 : ", num2)
	fmt.Println("ex 1 : ", num3)
	fmt.Println("ex 1 : ", num4)
	fmt.Println("ex 1 : ", num4-.1)
	fmt.Println("ex 1 : ", float32(num4-.1))
	fmt.Println("ex 1 : ", float64(num4-.1)) //부동소수점 오차 해결!
	fmt.Println("ex 1 : ", num5)
	fmt.Println("ex 1 : ", num6)
	fmt.Println("ex 1 : ", num7)
}
