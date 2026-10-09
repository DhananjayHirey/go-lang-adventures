package main

import (
	"fmt"
	"time"
)

func process(num int, c chan int) {
	time.Sleep(100 * time.Millisecond)
	c <- num + 1
}

func main() {
	num := 10
	c := make(chan int)
	go process(num, c)
	fmt.Print(num)
	result := <-c
	fmt.Print(result)

}
