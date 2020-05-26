package main

import (
	"fmt"
)

type dog struct {
	name   string
	weight int
}

type cat struct {
	name   string
	weight int
}

//struct dog method
func (d dog) bite() {
	fmt.Println(d.name, "a Dog bites!")
}

func (d dog) sounds() {
	fmt.Println(d.name, "a Dog barks!")
}

func (d dog) run() {
	fmt.Println(d.name, "a Dog is running!")
}

//struct cat method
func (d cat) bite() {
	fmt.Println(d.name, "a Cat bites!")
}

func (d cat) sounds() {
	fmt.Println(d.name, "a Cat cries!")
}

func (d cat) run() {
	fmt.Println(d.name, "a Cat is running!")
}

//animal behavior interface
type animal interface {
	bite()
	sounds()
	run()
}

//인터페이스를 파라메터로 받는다.
func act(animal animal) {
	animal.bite()
	animal.sounds()
	animal.run()
}

func main() {
	//interface 덕 타이핑
	//Go 덕 타이핑이란 : 메서드로만 타입을 판단한다.
	//구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 메서드로만 판단하는 방식
	//예를들어 오리처럼 날수 있고, 꽥꽥거리고, 뒤뚱뒤뚱 걷는 등 행동이 같으면 오리라고 볼 수 있다. -> 의미

	//exam 1
	dog := dog{"poll", 10}
	cat := cat{"bob", 5}

	//dog
	act(dog)
	//cat
	act(cat)
}
