 func maxSubArray(nums []int) int{
   maxCurrent, maxGlobal := nums[0], nums[0]
   for i := 1; i < len(nums); i++{
    maxCurrent = max(nums[i], maxCurrent+nums[i])
    if maxGlobal < maxCurrent{
      maxGlobal = maxCurrent
    }
   }
   return maxCurrent

 }

 fumc max(a, b int) int{
   if a > b{
     return a
   }
   return b
 }
