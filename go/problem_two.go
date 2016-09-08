// Find the sum of the even valued terms in the Fibonacci sequence whose values don't exceed 4 million.

package main

import(
  "fmt"
)

// fibSqnce collects terms in the Fibonacci sequence
// that are no greater than the 'max' parameter.
func fibSqnce(max int) []int {

  var fibs []int

  // Starting Numbers in Fibonacci Sequence
  mostRecent := 2
  secondMostRecent := 1

  // Get new Fibonacci Terms
  for i := 3; i <= max; i++ {

    if i == (mostRecent + secondMostRecent) && i < max {
      // Calculate next term in Fibonacci Sequence
      newFib := mostRecent + secondMostRecent

      // Replace previous Fibonacci terms
      secondMostRecent = mostRecent
      mostRecent = newFib

      // Place new Fib term in a slice
      fibs = append(fibs, newFib)
    }
  }
  fmt.Println("Terms:", fibs)
  return fibs
}

// Sum the numbers provided
func SumEvenVals(terms []int) int {
  var sum int
  for _, t := range terms {
    if t % 2 == 0 {
      fmt.Println(t)
      sum += t
    }
  }

  // Account for loop in FibSqnce beginning at 3
  sum += 2
  return sum
}

func main(){
  terms := fibSqnce(4000000)
  fmt.Println(SumEvenVals(terms))

}
