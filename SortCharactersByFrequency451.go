func frequencySort(s string) string {
  freq := make(map[rune]int)
  for _, char := range s {
    freq[char]++
  }
  type charFreq struct {
    char rune
    count int
  }
  
  var chars []charFreq
  for char, count := range freq {
    chars = append(chars, charFreq{char, count})
  }
  
  sort.Slice(chars, func(i,j int) bool {
    return chars[i].count > chars[j].count
  })
  
  var result strings.Builder
  for _, cf := range chars {
    result.WriteString(strings.Repeat(string(cf.char), cf.count))
  }
  return result.String()
}
