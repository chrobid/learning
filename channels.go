package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gen(c chan []int) {
	const listLength = 10000
	var num []int
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < listLength; i++ {
		num = append(num, rand.Intn(999999)) // generate random number here
		//fmt.Println(i)
	}
	c <- num
}

// concurrently generate lists of random numbers
func main() {
	listQuant := 10000
	c := make(chan []int)
	var nums [][]int

	for i := 0; i < listQuant; i++ {
		go gen(c)
	}
	for i := 0; i < listQuant; i++ {
		nums = append(nums, <-c)
	}
	fmt.Print("done")
}
