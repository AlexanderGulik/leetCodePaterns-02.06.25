func longestConsecutive(nums []int) int {
  setNums := make(map[int]bool)
  MaxStreak := 0 
  for _, num := range nums{
    setNums[num] = true
   }
  for num := range setNums{
    if !setNums[num-1] {
      currentNums := num
      currentStreak := 1

      for setNums[currentNums+1]{
        currentNums++
        currentStreak++
        
        
      }
      
        if currentStreak > MaxStreak {
          MaxStreak = currentStreak
        }
    }

  }
  return MaxStreak
  }

