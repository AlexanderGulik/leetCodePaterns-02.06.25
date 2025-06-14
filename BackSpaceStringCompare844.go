func backspaceCompare(s string, t string) bool {
  return processString(s) == processString(t)

}

func processString(str string) string {
  stack := []rune{}
  for _, char := range str {
    if char != '#' {
      stack = append(stack, char)
    } else id len(stack) > 0 {
      stack = stack[:len(stack)-1]
    }
  }
  return string(stack)
}
