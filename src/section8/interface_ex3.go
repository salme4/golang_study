package main

import (
	"fmt"
)

func main() {
	//실제 타입 검사 switch 문을 주로 사용
	//empty interface 는 Object 이므로, 타입 체크르 통해 형 변환 후 사용 가능

	//exam 1
	checkType(true)
	checkType(1)
	checkType(22.12)
	checkType(nil)
	checkType("hello golang!")
}

func checkType(args interface{}) {
	//args.(type)을 통해 현재 타입 반환
	switch args.(type) {
	case bool:
		fmt.Println("This is a bool", args)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int", args)
	case string:
		fmt.Println("This is a string", args)
	case float32, float64:
		fmt.Println("This is a float", args)
	default:
		fmt.Println("This is a nil", args)
	}
}
