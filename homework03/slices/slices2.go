package main

import "fmt"

func main() {

	a := []int64{1, 2, 5, 15}
	fmt.Println(reverse(a))
}

func reverse(a []int64) []int64 {
	reverseArray := make([]int64, len(a), cap(a))
	reverseArray = a[3:4]
	reverseArray = append(reverseArray, a[2], a[1], a[0])

	return reverseArray
}
