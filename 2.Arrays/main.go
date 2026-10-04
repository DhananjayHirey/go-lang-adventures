package main

import (
	"fmt"
)

func main() {
	// fmt.Println()

	var a [5]int
	fmt.Println("emp: ", a)
	a[4] = 100
	fmt.Println(a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	c := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(c)

	d := [...]int{200, 3: 400, 500, 600}
	fmt.Println(d)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	twoD2 := [3][4]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	fmt.Println(twoD2)
}
