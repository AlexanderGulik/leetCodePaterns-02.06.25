func partition(s string) [][]string{
  result := [][]string{}
  backtrack(&result, []string{}, s)
  return result
}

func backtrack(result *[][]string, current []string, s string){
  if len(s)==0 {
    temp := make([]string, len(current))
    copy(temp, current)
    *result = append(*result, temp)
    return
  }
  for i := 1; i < len(s); i++{
    prefix := s[:i]
    if isPalindrome(prefix){
      backtrack(result, append(current, prefix), s[i:])
    }
  }
}

func isPalindrome(s string) bool{
  left, right := 0, len(s) - 1
  for left < right{
    if s[left] != s[right]{
        return false
    }
    left++
    right--
  }
  return true
}
