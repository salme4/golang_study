package main

import (
	"fmt"
)

type dog struct {
	name   string
	weight int
}

type cat struct {
	name   string
	weight int
}

func printValue(s interface{}) {
	fmt.Println("ex 1 :", s)
}

func main() {
	//empty interface 활용
	//함수내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다. -> 모든 타입 지정 가능

	aDog := dog{"kuni", 4}
	aCat := cat{"bob", 3}

	printValue(aDog)
	printValue(aCat)
	printValue(10)
	printValue("animal")
	printValue(25.5)
	printValue([]dog{})
	printValue([5]dog{})
}
