func maxUniqueSplit(s string) int {
  max := 0 
  seen := make(map[string]bool)
  var backtrack func(int, int)
  backtrack = func(start, count int) {
    if start == len(s) {
      if count > max {
        max = count
      }
      return
    }
    for end := start + 1; end <= len(s); end++{
      substring := s[start:end]
      if !seen[substring] {
        seen[substring] = true
        backtrack(end, count+1)
        delete(seen, substring)
      }
    }
  }
  backtrack(0,0)
  return max
}


