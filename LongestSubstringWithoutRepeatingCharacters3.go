func lenghtOfLongestSubstring(s string) int {
  lastSeen := make(map[rune]inr)
  maxLength, left := 0,0
  for right, char := range s{
    if lastIndex, exiting := lastSeen[char]; exiting && lastIndex >= left{
      left = lastIndex+1
    }
    lastSeen[char] = right
    if right - left +1 > maxLength {
      maxLength = right - left +1
    }
  }
  return maxLength

}
