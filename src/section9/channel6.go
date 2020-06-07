package main

import (
	"fmt"
)

func main() {
	//채널의 반환 값 : (채널의 값, 수신여부)
	//close : 채널 닫기

	ch := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- "good!"
		}

	}()

	val1, ok1 := <-ch
	fmt.Println("exam 1 :", val1)
	fmt.Println("exam 1 :", ok1)

	val2, ok2 := <-ch
	fmt.Println("exam 1 :", val2)
	fmt.Println("exam 1 :", ok2)

	val3, ok3 := <-ch
	fmt.Println("exam 1 :", val3)
	fmt.Println("exam 1 :", ok3)

	close(ch)

	val4, ok4 := <-ch
	fmt.Println("exam 1 :", val4)
	fmt.Println("exam 1 :", ok4)
}
