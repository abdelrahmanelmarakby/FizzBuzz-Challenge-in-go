package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FIZZBUZZ  " + strconv.Itoa(i))
		} else if i%3 == 0 {
			fmt.Println("Fizz  " + strconv.Itoa(i))
		} else if i%5 == 0 {
			fmt.Println("Buzz  " + strconv.Itoa(i))
		}

	}
}
