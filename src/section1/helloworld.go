package main

import "fmt"

func main()  {
	/**
	* 실행방법 3가지
	* 1. go run [file] -> go file을 실행시킨다.
	* 2. go build [file] -> 실행 가능한 파일 생성한다.
	* 3. go install -> 실행 가능한 파일을 생성하여 GOBIN폴더로 이동한다.
	**/
	fmt.Println("hello")
	fmt.Println("world")
}