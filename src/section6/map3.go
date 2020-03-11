package main

import (
	"fmt"
)

func main11() {
	//Map
	//Map 값 변경 및 삭제

	//exam 1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
		"home1":  "http://test1.com",
	}

	fmt.Println(" ex 1 :", map1)

	map1["home2"] = "http://test2.com" //추가

	fmt.Println(" ex 1 :", map1)

	map1["home2"] = "http://test2-2.com" //수정

	fmt.Println(" ex 1 :", map1)
	fmt.Println()

	delete(map1, "home2") //삭제
	fmt.Println(" ex 1 :", map1)

	delete(map1, "google")
	fmt.Println(" ex 1 :", map1)

}
