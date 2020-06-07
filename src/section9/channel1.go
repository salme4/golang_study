package main

import (
	"fmt"
	"time"
)

func work1(v chan int) {
	fmt.Println("Work 1 : Start ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work 1 : End ---> ", time.Now())
	v <- 1
}

func work2(v chan int) {
	fmt.Println("Work 2 : Start ---> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Work 2 : End ---> ", time.Now())
	v <- 2
}

func main() {
	//채널
	//go routine 간의 상호 정보(데이터) 교환 및 실행 흐름 동기화 위해 사용 : 채널(동기식, 참조타입)
	//실행 흐름 제어 가능(동기, 비동기) -> 채널을 일반 변수로 선언 후 사용가능
	//데이터 전달 자료형 선언 후 사용 (지정 된 타입만 주고받을 수 있음)
	//interface{} 으로 자료형 상관없이 전달 가능.
	//값을 전달(복사 후 : bool, int 등), 포인터(슬라이스, 맵) 등을 전달 시에는 주의!! > 동기화 사용(Mutex)
	//멀티프로세싱 처리에서 교착상태(경쟁 : 데드락) 주의!!
	// <- , -> : 기호 사용 (채널 <- 데이터 : 송신), (변수 <- 채널 : 수신)

	//선언
	//var v chan int
	//v = make(chan int)
	v := make(chan int)

	//exam 1
	fmt.Println("Main Thread : Start ----> ", time.Now())
	go work1(v)
	go work2(v)

	//동기화 : time.Sleep 필요 없음.
	<-v
	<-v

	fmt.Println("Main Thread : End ----> ", time.Now())
}
