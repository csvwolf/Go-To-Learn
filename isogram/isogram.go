package isogram

import (
  "strings"
  "regexp"
)
func IsIsogram(input string) bool {
  m := make(map[rune] bool)
  r := regexp.MustCompile(`[^a-z]`)
  for _, v := range strings.ToLower(input) {
    if r.MatchString(string(v)) {
      continue
    }
    if m[v] {
      return false
    }
    m[v] = true
  }
  return true
}
