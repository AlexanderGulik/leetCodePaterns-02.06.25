func findMaxAverage(nums []int, k int) float64{
  if k > len(nums) || k < 0 {
    return 0
  }
  left, right := 0, 0
  subArr, result := 0, Math.MinInt32
  for right < len(nums){
    subArr += nums[right]
    if right >= k - 1 {
      if subArr > result {
        result = subArr
      }
      subArr -= nums[left]
      left++
    }
    right++
  }
  return float64(result)/float64(k)
}
