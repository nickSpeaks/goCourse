package main

import (
	"Golang/epam/goCourse/homework02/fibonacci"
	"fmt"
)

func main() {
	number := 16
	fmt.Printf("Hello, dear students! I would like to introduce the first %d numbers of the Fibonacci sequence: ", number)
	fibonacci.PrintNum(number)
}
