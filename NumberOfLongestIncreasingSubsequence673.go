func findNumberOfLIS(nums []int) int{
  lenghts := make([]int, len(nums))
  counts := make([]int, len(nums))
  for i := range lenghts {
    lenghts[i] = 1
    counts[i] = 1
  }

  for i := 1; i < len(nums); i++{
    for j := 0; j < i; j++{
      if nums[i] > nums[j]{
        if lenghts[j] + 1 > lenghts[i] {
            lenghts[i] = lenghts[j] + 1
            counts[i] = counts[j]
        } else if lenghts[j] + 1 == lenghts[i] {
          counts[i] += counts[j]
        }
      }
    } 
  }
  result, maxLen := 0, 0
  for _, i := range lenghts{
    if i > maxLen{
      maxLen = i
    }
  }
  for i := range counts{
    if maxLen == lenghts[i]{
      result += counts[i]
    }
  }
  return result


}
