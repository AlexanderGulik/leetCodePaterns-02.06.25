func minWindow(s string, t string) string {
  if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
    return ""
  }
  Window := make([]int, 128)
  count := len(t)
  left, right := 0,0
  minLen, startIdx := int(^uint(0)>>1), 0
  for _, char := range t {
    Window[char]++
  }
  for right < len(s) {
    if Window[s[right]] > 0 {
      count--
    }
    Window[s[right]]--
    right++
    for count == 0 {
      if right-left < minLen {
        startIdx = left
        minLen = right - left
      }
      if Window[s[left]] == 0{
        count++
      }
      Window[s[left]]++
      left++
    }
  }
  if minLen == int(^uint(0)>>1){
    return ""
  }
  return s[startIdx : startIdx+ minLen]
}
