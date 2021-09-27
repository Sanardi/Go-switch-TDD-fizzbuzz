package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Fizzbuzz prints and return an array of any size you pass to it/
//when the integer is divisible by 3 it prints "Fizz"
//When the int is divisible by 5 it prints "Buzz"
//when the int is divisible by 3 and 5 it prints "FizzBuzz"
func FizzBuzz(size int) []string {

	var output []string

	var wg sync.WaitGroup

	// adding 200 goroutines to the WaitGroup
	wg.Add(size + 1)

	for i := 1; i < size+1; i++ {
		ne := ""
		switch {
		case i%3 == 0 && i%5 == 0:
			ne = "FizzBuzz"
			wg.Done()
		case i%3 == 0:
			ne = "Fizz"
			wg.Done()
		case i%5 == 0:
			ne = "Buzz"
			wg.Done()
		default:
			ne = fmt.Sprint(i)
			wg.Done()
		}
		output = append(output, ne)
	}
	fmt.Println(output)
	return output
}

func main() {
	start := time.Now()
	FizzBuzz(1000)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
