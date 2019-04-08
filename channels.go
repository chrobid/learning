package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
func sum(nums []int, c chan int) {
	var sum int
	for _, value := range nums {
		sum += value
	}
	c <- sum
}
*/

func gen(c chan []int) {
	const listLength = 100
	var num []int
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < listLength; i++ {
		num = append(num, rand.Intn(999999)) // generate random number here
	}
	c <- num
}

// concurrently generate lists of random numbers
func main() {
	listQuant := 5
	c := make(chan []int)

	for i := 0; i < listQuant; i++ {
		go gen(c)
	}
	num1, num2, num3, num4, num5 := <-c, <-c, <-c, <-c, <-c
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)
}

// do we send the entire slice to the gen function, and then
// send that back to the channel? or do we send the individual
// values for each element in the slice to the channel, and
// then dump it from the channel into the slice?
//
// kk first it's not a slice it's an array and second we send the whole array
// back but don't have to send it there to begin with
