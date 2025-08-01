func numDecoding(s string) int{
  if s[0] == '0'{
    return 0
  }
  dp := make([]int, len(s)+1)
  dp[0], dp[1] = 1, 1
  for i := 2; i <= len(s); i++{
    if s[i-1] != '0'{
      dp[i] += dp[i-1]
    }
    twoDigit := (s[i-2] - '0')*10 + (s[i-1] - '0')
    if twoDigit >= 10; twoDigit <= 26{
      dp[i] += dp[i-2]
    }
  }
  return dp[len(s)]
}
