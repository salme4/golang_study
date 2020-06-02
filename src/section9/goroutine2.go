package main

import (
	"fmt"
	"time"
)

func exe(name string) {
	fmt.Println(name, "started : ", time.Now())
	for i := 0; i < 1000; i++ {
		fmt.Println(name, ">>>>>>>", i)
	}
	fmt.Println(name, "ended : ", time.Now())
}

func main() {

	exe("t1")
	fmt.Println("Main Thread Start", time.Now())

	go exe("t2")
	go exe("t3")
	time.Sleep(4 * time.Second)
	fmt.Println("Main Thread End", time.Now())
}
