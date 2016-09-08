// If we list all the natural numbers below 10 that are multiples of
// 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import(
  "fmt"
)

func findMultiples(num int, factors ...int) map[int]bool {

  // Gather all unique factors
  allMultiples := make(map[int]bool)

  for _, f := range factors {
    if num % f == 0 && allMultiples[f] == false {
      allMultiples[f] = true
    }
  }

  return allMultiples
}

func main(){
  fmt.Print(findMultiples(10, 3, 5))
}