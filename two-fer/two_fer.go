package twofer

func ShareWith(s string) string {
  if s != "" {
    return "One for " + s + ", one for me."
  }
  return "One for you, one for me."
}
