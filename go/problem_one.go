// If we list all the natural numbers below 10 that are multiples of
// 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import(
  "fmt"
)

func sum(nums []int) int {
  var sum int
  for _, v := range nums {
    sum += v
  }
  return sum
}

func main(){
  var multiples []int

  for i := 0; i < 1000; i++ {
    if i % 15 == 0 || i % 3 == 0 || i % 5 == 0 {
      multiples = append(multiples, i)
    }
  }

  fmt.Println(sum(multiples))
}