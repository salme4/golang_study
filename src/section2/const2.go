//상수 2
package main

import "fmt"

func main5() {
	//여러개 상수
	const a, b int = 10, 20
	const c, d, e = true, 0.2, "test"
	const (
		f bool  = false
		g int16 = 10
	)

	fmt.Println(a, b, c, d, e)
	fmt.Println(f, g)
}
