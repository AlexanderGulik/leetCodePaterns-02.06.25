func uniqueLetterString(s string) int {
  result, prev, total := 0,0,0
  hasMore := make([]int, 26)
  hasOne := make([]int, 26)
  for _, c := range s {
    idx := int(c - 'A')
    total++
    hasMore[idx] += hasOne[idx]
    prev += total - hasOne[idx] - hasMore[idx]
    hasOne[idx] = total - hasMore[idx]
    result += prev
  }

  return result
}
