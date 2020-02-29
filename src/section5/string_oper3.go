package main

import (
	"fmt"
	"strings"
)

func main() {
	//문자열 결합

	//exam 1 : 일반연산
	str1 := "The Go programming language is an open source project to make programmers more productive." +
		"Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that"

	str2 := "Instructions for downloading and installing the Go compilers, tools, and libraries."

	fmt.Println("ex 1: ", str1+str2)

	//exam 2 : Join
	strSet := []string{}
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)
	fmt.Println("ex 2 : ", strings.Join(strSet, " "))
}
