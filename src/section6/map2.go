package main

import (
	"fmt"
)

func main10() {
	//Map
	//Map 조회 및 순회(Iterator)

	//exam 1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}

	fmt.Println("ex 1 : ", map1["google"])
	fmt.Println("ex 1 : ", map1["daum"])
	fmt.Println()

	//exam 2 (순서가 없으므로 랜덤)
	for k, v := range map1 {
		fmt.Println("ex 2: ", k, v)
	}

	fmt.Println()

	for _, v := range map1 {
		fmt.Println("ex 2 : ", v)
	}

}
