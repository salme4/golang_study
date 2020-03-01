package main

import (
	"fmt"
)

func main1() {
	//배열
	//용량과 길이가 항상 같다.
	//배열 vs 슬라이스 차이점 중요
	//길이고정 vs 길이가변
	//값 타입 vs 참조 타입
	//복사 전달 vs 참조 값 전달
	//전체 비교연산자 사용 가능 vs 비교 연산자 사용불가
	//대부분 슬라이스 사용한다.

	//cap() : 배열, 슬라이스 용량
	//len() : 배열, 슬라이스 갯수

	//exam 1
	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}
	arr6 := [...]int{1, 2, 3, 4, 5} //배열 크기 자동 증가
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}

	arr1[2] = 5

	fmt.Printf("ex 1 : %-5T %d %v \n", arr1, len(arr1), arr1)
	fmt.Printf("ex 1 : %-5T %d %v \n", arr2, len(arr2), arr2)
	fmt.Printf("ex 1 : %-5T %d %v \n", arr3, len(arr3), arr3)
	fmt.Printf("ex 1 : %-5T %d %v \n", arr4, len(arr4), arr4)
	fmt.Printf("ex 1 : %-5T %d %v \n", arr5, len(arr5), arr5)
	fmt.Printf("ex 1 : %-5T %d %v \n", arr6, len(arr6), arr6)
	fmt.Printf("ex 1 : %-5T %d %v \n", arr7, len(arr7), arr7)

	//exam 2
	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}

	arr10 := [...]string{"kim", "lee", "park"}

	fmt.Printf("ex 2 : %-5T %d %v \n", arr8, len(arr8), arr8)
	fmt.Printf("ex 2 : %-5T %d %v \n", arr9, len(arr9), arr9)
	fmt.Printf("ex 2 : %-5T %d %v \n", arr10, len(arr10), arr10)
}
