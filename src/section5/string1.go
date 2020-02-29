package main

import (
	"fmt"
	"unicode/utf8"
)

func main9() {
	//**문자열**
	//큰 따옴표 "", 백스쿼트 `` 를 사용한다.
	//문자 char 존재 하지 않음
	//rune(int32) 문자 코드 값으로 표현
	//문자 : ''
	//자주 사용하는 escape : \\, \', \", \a(콘솔벨), \b(백스페이스), \f(쪽바꿈), \n(줄바꿈), \r(복귀), \t(탭)

	//exam 1
	var str1 string = "C:\\Users\\salme\\Document\\go_study" //escape
	str2 := `C:\Users\salme\Document\go_study`               //백스쿼트

	fmt.Println("ex 1 :", str1)
	fmt.Println("ex 1 :", str2)

	//exam 2
	var str3 string = "Hello, World!"
	var str4 string = "안녕하세요."

	fmt.Println("ex 2 : ", str3)
	fmt.Println("ex 2 : ", str4)

	fmt.Println()
	//exam 3 : len() -> 크기(byte)
	fmt.Println("ex 3 :", len(str3)) //space도 1byte
	fmt.Println("ex 3 :", len(str4))

	//exam 4 : 길이
	fmt.Println("ex 4 : ", utf8.RuneCountInString(str3)) // 좀 귀찮네..
	fmt.Println("ex 4 : ", utf8.RuneCountInString(str4))
	fmt.Println("ex 4 : ", len([]rune(str4))) //len사용한 길이 구하기.

}
