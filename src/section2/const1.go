//상수 1
package main

import "fmt"

func main4() {
	//상수
	//const 사용 초기화, 한 번 선언 후 값 변경 금지, 고정 된 값 관리용.
	const a string = "TEST1"
	const b = "TEST2"
	const c int32 = 10 * 10
	const d int = 22
	//const d = getHeight() 함수 사용 불가능
	const e = 35.6
	const f = false
	var result, err = fmt.Println("a : ", a)
	fmt.Println("result : ", result)
	fmt.Println("err : ", err)
}
