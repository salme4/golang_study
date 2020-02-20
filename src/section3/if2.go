package main

import "fmt"

func main2()  {
	//else 
	var a int = 50
	b := 70

	//exam 1
	if a >= 65 {
		fmt.Println("65 gt")
	} else {
		fmt.Println("65 lt")
	}

	if b >= 70 {
		fmt.Println("70 gt")
	} else {
		fmt.Println("70 lt")
	}

	//에러 발생 1 : else 줄 바꿈
	/*
		if b >= 70 {
			fmt.Println("70 gt")
		} else
		{
			fmt.Println("70 lt")
		}
	*/


}