package main

import (
	"fmt"
)

func rangeSum(n int, c chan int) {
	var sum int = 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	c <- sum
}

func main() {

	c := make(chan int)
	go rangeSum(1000, c)
	go rangeSum(7000, c)
	go rangeSum(5000, c)

	//순서대로 데이터 수신(동기) : 채널에서 값 수신 완료 될 때까지 대기.
	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println("exam 1 :", result1)
	fmt.Println("exam 1 :", result2)
	fmt.Println("exam 1 :", result3)
}
