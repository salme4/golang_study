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
func (d dog) run() {
	fmt.Println(d.name, "a Dog is running!")
}

//struct cat method
func (d cat) run() {
	fmt.Println(d.name, "a Cat is running!")
}

//익명 interface
func act(ani interface{ run() }) {
	ani.run()
}

func main() {
	//익명 interface : 즉시 선언 후 사용 가능.

	//exam 1
	dog := dog{"poll", 10}
	cat := cat{"bob", 5}

	//dog
	act(dog)
	//cat
	act(cat)

}
