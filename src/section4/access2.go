package main

import (
	"fmt"
	check "section4/lib"
	_ "section4/lib2"
)

func main4() {
	//별칭 사용 : alias
	//빈 식별자 사용 : _

	fmt.Println("10보다 큰수 ? : ", check.CheckNum(10))
}
