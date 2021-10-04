package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

// Fizzbuzz prints and return an array of any size you pass to it/
//when the integer is divisible by 3 it prints "Fizz"
//When the int is divisible by 5 it prints "Buzz"
//when the int is divisible by 3 and 5 it prints "FizzBuzz"

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
	var wg sync.WaitGroup
	//var output []string
	output := make([]string, size)
	//c1 := make(chan string, size)
	c1 := make(chan string, 1)
	//fmt.Println("Channel's capacity:", cap(c1)) // => size
	wg.Add(size)

	for i := 1; i < size+1; i++ {
		go func(i int, wg *sync.WaitGroup) {
			//defer wg.Done()
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
			c1 <- ne

		}(i, &wg)

		output[i-1] = <-c1
		//fmt.Println("output: ", output)
		wg.Done()
	}

	wg.Wait()
	close(c1)
	fmt.Println("output: ", output)
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)

}
