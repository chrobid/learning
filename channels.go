package main

func sum(nums []int, c chan int) {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	c <- sum
}

func gen() int {
	return 5
}

func main() {
	length := 100
	var s []int
	for i, _ := range s {
		s[i] = gen()
	}
}
