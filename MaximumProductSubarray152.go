func maxProduct(nums int[]) int{
  minCurrent := nums[i]
  maxCurrent := nums[0]
  maxGlobal := nums[0]
  for i := 1; i< len(nums); i++{
    prevMax := maxCurrent
    prevMin := minCurrent
    maxCurrent = max(nums[i], max(prevMin*nums[i], prevMax * nums[i]))
    minCurrent = min(nums[i], min(prevMin*nums[i], prevMax * nums[i]))
    if maxCurrent > maxGlobal{
      maxGlobal = maxCurrent
    }
  }
  return maxGlobal

}

func max (a,b int) int{
 if a > b{
   return a
 }
 return b
}


func min (a,b int) int{
 if a < b{
   return a
 }
 return b
}


