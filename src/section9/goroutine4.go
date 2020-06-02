package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//go routine closure 사용

	//exam 1
	runtime.GOMAXPROCS(1)

	s := "goRoutine Closure : "

	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) //반복문 클로져는 일반적으로 즉시 실행 but, 고루틴 클로져는 가장 나중에 실행(반복문 종료 후 실행된다.)
	}

	time.Sleep(3 * time.Second)
}
