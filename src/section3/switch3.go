package main

import "fmt"

func main()  {
	//다중 조건
	switch a:= 30; a / 15 {
	case 2, 4, 6:
		fmt.Println("a / 15 ", a, "는 짝수")
	case 1,3,5:
		fmt.Println("a / 15", a, "는 홀수")
	}

	//fallthrough : 통과 시켜 실행한다., 마지막 case 에 사용 시 에러, 의존 조건일 경우 사용 가능
	switch e := "go"; e {
	case "java":
		fmt.Println("java")
	case "go":
		fmt.Println("go")
		fallthrough
	case "python":
		fmt.Println("python")
	case "ruby":
		fmt.Println("ruby")
		fallthrough
	case "c++":
		fmt.Println("c++")
	}
}