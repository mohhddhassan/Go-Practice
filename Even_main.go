package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			numbers <- i
		}
		close(numbers) 
	}()
	for num := range numbers {
		if num%2 == 0 {
			fmt.Printf("Even number received: %d\n", num)
		}
	}
}

