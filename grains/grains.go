package grains

import "errors"

func Square(input int) (uint64, error) {
  result := uint64(1)
  if input <= 0 || input > 64 {
    return result, errors.New("error")
  }
  for i := 0; i < input - 1; i++ {
    result *= 2
  }
  return result, nil
}

func Total() uint64 {
  result := uint64(0)
  for i := 1; i <= 64; i++ {
    tmp, _ := Square(i)
    result += tmp
  }
  return result
}
