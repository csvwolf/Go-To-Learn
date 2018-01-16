package luhn

import (
  "regexp"
  "strings"
  "strconv"
)

func checkInput(input string) bool {
  r := regexp.MustCompile(`^[0-9]+$`)
  return len(input) > 1 && r.MatchString(input)
}


func Valid(input string) bool {
  result := 0
  str := strings.Join(strings.Split(input, " "), "")
  if !checkInput(str) {
    return false
  }
  for i := 0; i < len(str); i++ {
    tmp, _ := strconv.Atoi(string(str[len(str) - i - 1]))
    if i % 2 != 0 {
      if tmp * 2 > 9 {
        result += tmp * 2 - 9
      } else {
        result += tmp * 2
      }
    } else {
      result += tmp
    }
  }

  return result % 10 == 0
}
