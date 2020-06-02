package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func exeE(name int) {
	r := rand.Intn(100)
	fmt.Println(name, "start :", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>>", r, i)
	}
	fmt.Println(name, "end :", time.Now())
}

func main() {
	//멀티 코어 ( 다중 CPU ) 최대한 활용

	runtime.GOMAXPROCS(runtime.NumCPU())                       //현 시스템 CPU 코어 갯수 반환 후 설정
	fmt.Println("Current System CPU :", runtime.GOMAXPROCS(0)) //설정 값 출력

	//exam 1
	fmt.Println("Main Thread Start :", time.Now())
	for i := 0; i < 100; i++ {
		go exeE(i) //go routine 100 생성
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Main Thread End :", time.Now())
	fmt.Println("Current System CPU :", runtime.GOMAXPROCS(0)) //설정 값 출력
}
