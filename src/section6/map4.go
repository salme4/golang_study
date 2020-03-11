package main

import (
	"fmt"
)

func main12() {
	//Map : 주의사항
	//Map 조회 시 주의사항

	//exam 1
	map1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1, lemonExists := map1["lemon"]
	value2, exists := map1["kiwi"] //없는 값은 데이터 타입의 초기화 값이 리턴된다.
	value3, ok := map1["kiwi"]     //리턴 값이 두개이다. 두 번째는 키의 존재 유무

	fmt.Println("ex 1 : ", value1, lemonExists)
	fmt.Println("ex 1 : ", value2, exists)
	fmt.Println("ex 1 : ", value3, ok)

	//exam 2
	if value, isExists := map1["kiwi"]; isExists {
		fmt.Println("ex 2 : ", value)
	} else {
		fmt.Println("kiwi is not exists")
	}

	if value, isExists := map1["banana"]; isExists {
		fmt.Println("ex 2 : ", value)
	} else {
		fmt.Println("banana is not exists")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("kiwi is not exists")
	}
}
