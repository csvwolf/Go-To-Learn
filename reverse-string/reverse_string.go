package reverse

func String(input string) string {
  c := []rune(input)
  result := make([]rune, len(c))
  for i := len(c) - 1; i >= 0; i-- {
    result[len(c) - 1 - i] = c[i]
  }
  return string(result)
}
