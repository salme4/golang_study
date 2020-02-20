//변수2
package main

import "fmt"

func main2()  {
	//여러 개 선언
	var (
		name string = "dave"
		height int 
		weight float32 
		isRunning bool
	)

	height = 173
	weight = 82.3
	isRunning = true

	fmt.Println("name : ", name, " height : ", height, " weight : ", weight, " isRunning : ", isRunning)
}