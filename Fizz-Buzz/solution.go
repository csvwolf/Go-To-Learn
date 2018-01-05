package main

import "fmt"

// Using go build . to build the project

func FizzBuzz(n int) []string {
  fizz := make([]string, n)
  for i := 1; i <=n; i++ {
    if i % 3 == 0 {
      fizz[i -1] = "Fizz"
    }
    if i % 5 == 0 {
      fizz[i - 1] += "Buzz"
    }
    if fizz[i - 1] == "" {
      fizz[i - 1] = fmt.Sprintf("%d", i)
    }
  }

  return fizz
}
