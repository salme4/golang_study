package main

import (
	"fmt"
	"runtime"
)

func main() {
	//채널의 버퍼 사용 (비동기)

	runtime.GOMAXPROCS(4)
	//exam 1 : with buffer
	ch := make(chan bool, 4)
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go Routine : ", i)
		}
	}()

	//buffer : (송신)데이터를 가득 채우고 대기, (수신) 버퍼가 비워질 때까지 가져옴.
	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main :", i)
	}

}
