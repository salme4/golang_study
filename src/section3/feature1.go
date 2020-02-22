package main

import "fmt"

func main11() {
	//모호하거나 애매한 문법을 금지
	//후치 연산 : 반환값이 없음. (i++), 전치 연산은 허용하지 않음(++i X)
	//증감연산 반환값 없음 : sum += i++ X
	//포인터 허용, 연산은 비허용

	//exam 1
	sum, i := 0, 0
	for i <= 100 {
		//sum += i++ : Error!
		sum += i
		i++
		//++i : Error!
		fmt.Println("ex 1:", sum)
	}

}
