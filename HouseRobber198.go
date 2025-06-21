func rob(nums []int) int{
  currentMax, prevMax := 0,0
  for _, num := range nums{
    temp := currentMax
    currentMax = max(prevMax+num, currentMax)
    prevMax = temp
  }
  return currentMax
}

func mac(a, b int) int{
  if a > b {
    return a
  }
  return b
}


