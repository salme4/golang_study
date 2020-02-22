package main

import "fmt"

func main9()  {
	//exam 1 : 
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i 
	}
	fmt.Println("ex 1 : sum", sum)

	//exam 2
	var sum2, i int
	for i <= 100 {
		sum2 += i
		i++
		// j := i++ go에서 후치연산은 반환값이 없다. X
	}
	fmt.Println("ex 2 : sum2", sum2)

	//exam 3
	sum3, j := 0, 0
	for {
		if j > 100 {
			break
		}
		sum3 += j
		j++ 
	}
	fmt.Println("ex 3 : sum3", sum3)


	//exam 4
	for k, j := 0, 0; k <= 10; k, j = k + 1, j+10 {
		fmt.Println("ex 4", k, j)
	}

	//에러 발생
	/*
	for k, j := 0, 0; k <= 10; k++, j += 10 { }
	*/


}