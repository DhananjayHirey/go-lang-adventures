package main

import (
	"fmt"
)

func main() {
	// var age int = 32
	// fmt.Println(age)

	// i := 0
	// for i <= 3 {
	// 	fmt.Print(i)
	// 	i += 1
	// }
	// fmt.Println()
	// for j := 0; j <= 3; j++ {
	// 	fmt.Print(j)
	// }
	// fmt.Println()
	// for k := range 3 {
	// 	fmt.Println("range", k)
	// }

	// for {
	// 	fmt.Println("loop")
	// 	break
	// }

	// for n := range 6 {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	// if num := 9; num < 0 {
	// 	fmt.Println(num, "is negative")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println(time.Now().Day(), " : It's a weekend")
	// default:
	// 	fmt.Println(time.Now().Day(), " : It's a weekday!")
	// }

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	//     fmt.Println("It's before noon")
	// default:
	//     fmt.Println("It's after noon")
	// }

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}
	whatAmI(3)
	whatAmI(true)
	whatAmI("Hello")

}
