func leastInterval(tasks []byte, n int) int {
  freq := make([]int, 26)
  for _, task := range tasks {
    freq[task-'A']++
  }
  maxFreq := 0
  countMaxFreq := 0
  for _, f := range freq {
    if f > maxFreq {
      maxFreq = f
      countMaxFreq = 1
    } else if f == maxFreq {
      countMaxFreq++
    }
  }
  intervals := (maxFreq - 1) * (n + 1)+ countMaxFreq
  if intervals < len(tasks){
      return len(tasks)
  }
  return intervals
}
