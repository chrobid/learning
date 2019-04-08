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
		fmt.Println(i)
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
