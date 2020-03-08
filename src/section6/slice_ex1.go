package main

import "fmt"

func main6()  {
	//슬라이스 추가 및 병합

	//exam 1
	s1 := []int{1, 2, 3, 4, 5} //길이와 크기 : 5, 10
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	s1 = append(s1, 6, 7)  

	s2 = append(s1, s2...) //슬라이스를 삽입 할 경우 ... 사용
	s3 = append(s2, s3[0:3]...) //추출 후 병합

	fmt.Println("ex 1 : ", s1)
	fmt.Println("ex 1 : ", s2)
	fmt.Println("ex 1 : ", s3)

	//exam 2 : 용량(크기) 증가 방식 
	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("ex 2 -> len : %d, cap : %d, value : %v\n", len(s4), cap(s4), s4 ) //길이및 용량 자동 증가
	}

	
}