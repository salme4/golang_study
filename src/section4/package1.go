package main

import (
	"fmt"
	"os"
)

func main1() {
	//패키지 : 코드 모듈화 및 재사용
	//응집도 높이고, 결합도는 낮추고
	//Go : 패키지 단위의 독립적이고 작은 단위로 개발 -> 작은 패키지를 결합해서 작성 할 것을 권고
	//패키지 이름 == 디렉토리 이름
	//같은 패키지 내 -> 소스파일들은 디렉토리명을 패키지 명으로 사용한다.
	//***네이밍 규칙 : 소문자 private, 대문자 public***
	//Go : main 패키지는 특별하게 인식 -> 컴파일러 공유 라이브러리 X, 프로그램의 시작점 entryPoint

	var name string

	fmt.Println("이름은? : ")

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi %s\n", name)
}
