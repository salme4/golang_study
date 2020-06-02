package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 func end", time.Now())
}

func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 func end", time.Now())
}

func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 func end", time.Now())
}

func main() {
	//go routine
	//다른 언어의 thread 와 비슷한 기능
	//생성 방법 간단하며, 리소스 매우 적게 사용
	//즉, 수 많은 고루틴 동시 생성 실행 가능
	//비 동기적 함수 루틴 실행 -> 채널을 통한 통신 가능
	//공유 메모리 사용 시에 정확한 동기화 코딩 필요
	//싱글루틴에 비해 항상 빠르지는 않다.

	//멀티 쓰레드 장점과 단점
	//장점 : 응답성 향상, 자원공유를 효율적으로 활용 가능, 작업이 분리되어 코드 간결
	//단점 : 구현하기 어려움, 테스트 및 디버깅 어려움, 전체 프로세스에 사이드이펙트, 성능 저하, 동기화 코딩 반드시 숙지, 데드락...

	exe1()

	fmt.Println("Main Thread Start", time.Now())

	go exe2()
	go exe3()
	time.Sleep(4 * time.Second) //main thread가 종료되면 다른 sub thread들도 종료된다. ( 데몬 쓰레드 )
	fmt.Println("Main Thread End", time.Now())
}
