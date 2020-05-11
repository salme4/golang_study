package main

import (
	"fmt"
)

type cart struct {cnt, price int}

//결제 함수
func (b cart) purchase() int {
	return b.cnt * b.price
}

//원본 수정 (참초 전달 형식)
func (b *cart) reOrderR(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

//원본 수정 X (값 전달 형식)
func (b cart) reOrderV(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main4() {
	//receiver(구조체 메서드) 전달(값, 참조) 형식
	//함수는 기본적으로 값 호출 -> 변수의 값이 복사 후 내부 전달(원본 수정X) -> 맵, 슬라이스 등은 참조 전달
	//receiver(구조체)도 마찬가지로 포인터를 활용해서 메서드 내에서 원본 수정 가능
	
	//exam 1
	cart1 := cart{3, 5000}
	fmt.Println("ex 1 (totPrice) :", cart1.purchase())

	cart1.reOrderR(7, 5000)  //원본 변경 (참조 전달)
	fmt.Println("ex 2 (totPrice) :", cart1.purchase())

	cart1.reOrderV(10, 0)  //원본 변경 X (값 전달)
	fmt.Println("ex 3 (totPrice) :", cart1.purchase())
}