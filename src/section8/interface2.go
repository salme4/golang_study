package main

import (
	"fmt"
)

type dog struct {
	name   string
	weight int
}

func (d dog) bite() {
	fmt.Println(d.name, "bite!")
}

//동물의 행동 인터페이스
type behavior interface {
	bite()
}

func main() {

	//interface 구현 예제
	//exam 1
	dog1 := dog{"kuni", 10}
	var interface1 behavior

	//싱기방기
	interface1 = dog1
	interface1.bite()

	//exam 2
	dog2 := dog{"marry", 20}
	inter2 := behavior(dog2)
	inter2.bite()

	//exam 3 : 슬라이스
	inters := []behavior{dog1, dog2}
	for idx, _ := range inters {
		inters[idx].bite()
	}

	//값 형태로 실행
	for _, val := range inters {
		val.bite()
	}

}
