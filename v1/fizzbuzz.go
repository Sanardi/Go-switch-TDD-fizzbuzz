package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Fizzbuzz prints and return an array of any size you pass to it/
//when the integer is divisible by 3 it prints "Fizz"
//When the int is divisible by 5 it prints "Buzz"
//when the int is divisible by 3 and 5 it prints "FizzBuzz"
func FizzBuzz(size int) []string {
	var output []string
	for i := 1; i < size+1; i++ {
		ne := ""
		switch {
		case i%3 == 0 && i%5 == 0:
			ne = "FizzBuzz"
		case i%3 == 0:
			ne = "Fizz"
		case i%5 == 0:
			ne = "Buzz"
		default:
			ne = fmt.Sprint(i)
		}
		output = append(output, ne)
	}
	fmt.Println(output)
	return output
}

func main() {
	start := time.Now()

	s := os.Args[1]
	//flag.Parse()
	size, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(size)

	fmt.Printf("size: %d\n", size)
	FizzBuzz(size)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
