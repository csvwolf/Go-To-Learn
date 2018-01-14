package diffsquares

func SquareOfSums(n int) int {
  result := 0
  for i := 1; i <=n; i++ {
    result += i
  }
  return result * result
}

func SumOfSquares(n int) int {
  result := 0
  for i := 1; i <=n; i++ {
    result += i*i
  }
  return result
}

func Difference(n int) int {
  return SquareOfSums(n) - SumOfSquares(n)
}
