package main

import (
	"fmt"
	"time"
)

func main() {
	//채널의 버퍼 미사용

	//exam 1 : without buffer
	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go Routine : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	//채널 사용 시 송신 후 수신할 때까지 동기식으로 기다린다.
	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main :", i)
	}

}
