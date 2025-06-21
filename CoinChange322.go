func coinChange(coins []int, amount int) int{
  dp := make([]int, amount+1)
  for i := 1; i <= amount; i++{
    dp[i] = math.MaxInt32
    for _, coin := range coins {
      if i >= coin && dp[i-coin] != math.MaxInt32{
        dp[i] = min(dp[i], dp[i-coin]+1)
      }
    }
  }
  if dp[amount] == math.MaxInt32{
    return -1
  }
  return dp[amount]

}

func min(a, b int) int{
  if a < b{
    return a
  }
  return b
}
