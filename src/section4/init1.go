package main

import (
	"fmt"
)

func init() {
	//패키지 로드시에 가장 먼저 호출 됨.
	//가장 먼저 초기화 되는 작업 작성 시 유용하다.
	fmt.Println("main 보다 먼저 실행됨.")
}

func main5() {
	fmt.Println("main start")
}
