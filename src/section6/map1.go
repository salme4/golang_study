package main

import (
	"fmt"
)

func main9() {
	//Map
	//hashTable, dictionary, Key-Value로 자료 저장
	//참조형 타입 : 참조 값 전달
	//비교연사자 사용 불가능 (참조타입이라서)
	//특징 : 참조타입(Key)로 사용 불가능(Key에 참조타입 넣을 수 없음), 값(value)으로 모든 타입 사용가능
	//make() 함수 및 축약(리터럴)로 초기화 가능
	//순서가 없음.

	//exam 1
	var map1 map[string]int = make(map[string]int) //정석 선언 및 초기화
	var map2 = make(map[string]int)                // 자료형 생략
	map3 := make(map[string]int)                   // 축약형 선언 및 초기화

	fmt.Println("ex 1 : ", map1)
	fmt.Println("ex 1 : ", map2)
	fmt.Println("ex 1 : ", map3)

	//exam 2
	map4 := map[string]int{} //json 형태이다.
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	map5 := map[string]int{
		"apple":  15,
		"banana": 40,
		"orange": 23,
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	fmt.Println("ex 2 :", map4)
	fmt.Println("ex 2 :", map5)
	fmt.Println("ex 2 :", map6)
	fmt.Println("ex 2 :", map6["orange"])

}
