func longestPalidrome(s string) string{
  check := func(i, j int) bool{
    left, right := i, j-1
    for left < right{
      if s[left] != s[right]{
        return false
      }
      left++
      right--
    }
    return true
  }

  for longest := len(s); longest > 0; longest--{
    for start := 0; start <= len(s)-longest; start++{
      if check(start, start+longest){
        return s[start : start+longest]
      }
    }
  }
  return ""

}
