// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"
import "sort"

// findFactors Gathers All Factors
func findFactors(n int) []int {
  
  var factors []int

  for i := 1; i < n; i++ {
    if n % i == 0 {
      n = n/i
      factors = append(factors, i)
      factors = append(factors, n)
    }
  }
  return factors
}

// isPrime...
func isPrime(n int) bool {
  return len(findFactors(n)) == 2
}

// Largest...
func Largest(nums sort.IntSlice) int {
  largest := nums[0]

  for _, n := range nums {
    if n > largest {
      largest = n
    }
  }

  return largest
}

func main() {
  // Container for Prime Factors
  var primeFactors []int

  // Gather All Factors
  allFactors := findFactors(600851475143)

  // Select Prime Factors
  for _, f := range allFactors {
    if isPrime(f) {
      primeFactors = append(primeFactors, f)
    }
  }

  // Sort and Return Largest Prime Factor
  fmt.Println("Largest prime factor: ", Largest(primeFactors))
}
