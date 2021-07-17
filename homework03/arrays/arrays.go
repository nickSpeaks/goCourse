package main

import "fmt"

func main() {
    nums := []int{1, 2, 3, 4, 5, 6}
    average(nums...)
}

func average(nums ...int) {

    sum := 0
    for _, num := range nums {
        sum += num
    }
    s := float64(sum)
    l := float64(len(nums))
    fmt.Println(s / l)
}