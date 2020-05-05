package main

import (
	"fmt"
)

func main13() {
	//closure
	//익명함수 사용할 경우 함수를 변수에 할당해서 사용가능
	//함수 안에서 함수를 선언 및 정의 가능 -> 외부 함수에 선언 된 변수에 접근 가능한 함수
	//함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
	//closure 사용 이유 : 함수를 호출할 때 이전에 존재했던 값을 유지하기 위해서..  -> 비동기, 누적 카운트, 무분별한 전역변수 남발을 자제
	//남발 -> 객체들이 메모리 존재하므로 -> 메모리 부족, 오버플로우 발생

	//ex 1 : closure 사용안함
	multiflyFun := func(x, y int) int {
		return x * y
	} //익명 함수 할당

	r1 := multiflyFun(5, 10)

	fmt.Println("ex 1 :", r1)

	//ex 2 : closure 사용
	m, n := 5, 10               //변수가 캡쳐
	sumFun := func(c int) int { //익명함수에 변수 할당
		return m + n + c //지역 변수가 소멸되지 않는다. (함수 호출 시 마다 사용 가능)
	}

	r2 := sumFun(10)

	fmt.Println("ex 2 :", r2)
}
