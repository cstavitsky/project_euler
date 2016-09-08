// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"
import "sort"

func findFactors(n int) []int {
  
  var factors []int

  // Collect Factors
  // Begin at n/47 because of number's huge size
  if n == 600851475143 {
    for i := n/47; i > 0 ; i-- {
      if i % 2 == 0 || i % 3 == 0 || i % 5 == 0 || i % 7 == 0 || i % 11 == 0 || i % 13 == 0 || i % 17 == 0 || i % 19 == 0 || i % 23 == 0 || i % 29 == 0 {
        continue
      }
      // If n is divisible by i, i is a Factor
      if n % i == 0 && n > 1 {
        factors = append(factors, i)
      } else {
        continue
      }
    }
  } else {
    // for smaller numbers
    for i := n; i > 0 ; i-- {
      if i % 2 == 0 || i % 3 == 0 || i % 5 == 0 || i % 7 == 0 || i % 11 == 0 || i % 13 == 0 || i % 17 == 0 || i % 19 == 0 || i % 23 == 0 || i % 29 == 0 {
        continue
      }
      // If n is divisible by i, i is a Factor
      if n % i == 0 && n > 1 {
        factors = append(factors, i)
      } else {
        continue
      }
    }
  }

  return factors
}

func isPrime(n int) bool {
  if len(findFactors(n)) == 2 {
    fmt.Print(n, "IS PRIME")
    return true
  }
  return false
}

func Largest(nums sort.IntSlice) int {
  // sort.Sort(sort.Reverse(nums))
  return nums[0]
}

func main() {
  
  var primeFactors []int

  factors := findFactors(600851475143)

  fmt.Println("Finding factors...")
  for _, f := range factors {
    if isPrime(f) {
      primeFactors = append(primeFactors, f)
    }
  }

  // largest := Largest(primeFactors)
  fmt.Println("Prime Factors: ", primeFactors)
  fmt.Println("Solution", primeFactors[0])
}