func rob (nums []int) int{
  if len(nums) == 1{
    return nums[0]
  }
  return max(robHelp(nums[1:]), robHelp(nums[:len(nums)-1]))
}

func robHelp( nums []int) int{
  prevprev, prev := 0,0
  for _, num := range nums{
    current := max(prevprev+num, prev)
    prevprev, prev = prev, current
  }
  return prev
}

func max (a, b int) int{
  if a > b{
    return a
  }
  return b
}
