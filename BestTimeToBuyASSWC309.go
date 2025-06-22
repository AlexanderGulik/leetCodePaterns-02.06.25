func maxProfit(prices []int) int{
  hold := -prices[0]
  sold := math.MinInt32
  rest := 0
  for i := 1; i < len(prices); i++{
    prevHold := hold
    hold = max(hold, rest-prices[i])
    rest = max(rest, sold)
    sold = prevHold + prices[i]
  }
  return max(rest,sold)
}

func max(a, b int) int{
  if a > b {
    return a
  }
  return b
}
