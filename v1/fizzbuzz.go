package main

import "fmt"

// Sum calculates the total from a slice of numbers.
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
