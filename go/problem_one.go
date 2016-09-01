package main 

import("fmt")

//Find the sum of all multiples of 3 and 5 below 1000

func Add(nums []int) int {
	var sum int
	for _, i := range nums {
		sum += i
	}
	return sum
}

func FindMultiples() []int {
	var multiples []int
	for i := 0; i < 1000; i++ {
		if i % 15 == 0 || i % 5 == 0 || i % 3 == 0 {
			multiples = append(multiples, i)
		}
	}
	return multiples
}

func main(){
	sum := Add(FindMultiples())
	fmt.Println(sum)
}