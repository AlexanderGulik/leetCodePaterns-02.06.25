func canJump(nums []int) bool{
  maxReach := 0
  for idx := 0; idx < len(nums); idx++{
    if maxReach < idx {
      return false
    }
    if idx + nums[idx] > maxReach {
      maxReach = idx + nums[idx]
    }
  }
  return true
}
