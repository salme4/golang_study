package main

import (
	"fmt"
)

type car struct {
	color string
	name  string
}

func main7() {
	c1 := car{"red", "220d"}
	c2 := new(car)
	c2.color, c2.name = "white", "sonata"
	c3 := &car{}
	c4 := &car{"black", "520d"}

	fmt.Println("ex 1: ", c1)
	fmt.Println("ex 1: ", c2)
	fmt.Println("ex 1: ", c3)
	fmt.Println("ex 1: ", c4)

}
