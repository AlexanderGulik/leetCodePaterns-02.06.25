func canPartition(nums int) bool {
  totalSum := 0
  for _, num := range nums {
    totalSum += num
  }
  if totalSum % 2 != 0 {
    return false
  }
  target := totalSum / 2
  dp := make([]bool, target + 1)
  for _, num := range nums {
    for i := target; i >= num; i--{
        if dp[i - num]{
          dp[i] = true
        }
    }
  }
  return dp[target]         
}
