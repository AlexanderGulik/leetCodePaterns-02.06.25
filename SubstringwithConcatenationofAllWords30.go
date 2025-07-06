func findSubstring(s string, words []string) []int {
  if len(words) == 0 || len(s) == 0 {
    return []int{}
  }
  wordLength := len(words[0])
  wordCount := len(words)
  result := []int{}
  wordFreq := make(map[string]int)
  for _, word := range words {
    wordFreq[word]++
  }
  for i := 0; i < wordLength; i++{
    left, right := i, i
    seenWords := make(map[string]int)
    count := 0
    for right+wordLength <= len(s) {
      word := s[right : right+wordLength]
      right += wordLength
      if _, exists := wordFreq[word]; exists {
        seenWords[word]++
        count++
        for seenWords[word] > wordFreq[word] {
          leftWord := s[left : left+wordLength]
          seenWords[leftWord]--
          count--
          left += wordLength
        }
        if count == wordCount {
          result = append(result, left)
        }
      } else {
        seenWords = make(map[string]int)
        count = 0
        left = right
      }
    }
  }
  return result
}
