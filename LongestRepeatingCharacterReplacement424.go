func characterReplacement(s string, k int) int{
  count := [26]int{}
  maxCount :=0
  maxLenght :=0
  left := 0
  for right := 0; right < len(s); right++{
    count[s[right]-'A']++
    if count[s[right]-'A'] > maxCount {
      maxCount = count[s[right]-'A']
    }
    if (right-left+1-maxCount > k {
      count[s[left]-'A']--
      left++
    }
    if right-left+1 > maxLenght{
      maxLenght = right-left+1
    }
  }
  return maxLenght
}
