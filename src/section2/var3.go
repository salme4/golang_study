//변수 3
package main

import "fmt"

func main3()  {
	//짧은 선언
	//함수 안에서만 사용한다, 선언 후 할당 하면 예외 발생
	//주로 제한 된 범위의 함수내에서 사용 할 경우 코드 가독성을 높일 수 있다.
	
	shortVar1 := 3
	shortVar2 := "test"
	shortVar3 := false

	// shortVar := true -> 예외발생

	fmt.Println("shortvar 1 : ", shortVar1)
	fmt.Println("shortvar 2 : ", shortVar2)
	fmt.Println("shortvar 3 : ", shortVar3)

	//example
	if i:=1; i < 15 {
		fmt.Println("short variable test success")
	}
}