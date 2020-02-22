package main

import "fmt"

func main()  {
	//exam 1 : 반복문 Loop label 사용
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("ex 1 :", i, j)
		}
	}

	//exam 2 : continue 활용 홀 수만 출력
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("ex 2 : ", i)
	}

	//exam 3 : continue Loop label 사용
Loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("ex 3 : ", i, j)
		}
	}
}