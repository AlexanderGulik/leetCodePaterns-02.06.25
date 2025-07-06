func longestWord(words []string) string {
  sort.Strings(words)
  validWords := make(map[string]bool)
  result := ""
  for _, word := range words {
    if len(word) == 1 || validWords[word[len(word)-1]] {
      validWords[word] = true
      if len(result) < len(word){
          result = word
      } else if len(word) == len(result) && result < word {
        result = word
      }
    }
  }
  return result
}
