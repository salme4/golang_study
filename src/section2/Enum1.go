//열거형 1
package main

import "fmt"

func main6() {
	//열거형
	//상수를 사용하는 일정한 규칙에 따라 숫자를 계산 및 증가시키는 묶음.
	const (
		JAN = 1
		FEB = 2
		MAR = 3
		APR = 4
		MAY = 5
		JUN = 6
	)

	fmt.Println(JAN)
	fmt.Println(FEB)
	fmt.Println(MAR)
	fmt.Println(APR)
	fmt.Println(MAY)
	fmt.Println(JUN)
}
