package main

import (
	"fmt"
	"reflect"
)

type Car2 struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
}

func main9() {
	//구조체 필드에 태그를 사용

	//exam 1
	tags := reflect.TypeOf(Car2{})

	for i := 0; i < tags.NumField(); i++ {
		fmt.Println("ex 1 :", tags.Field(i).Tag, tags.Field(i).Name, tags.Field(i).Type)
	}
}
