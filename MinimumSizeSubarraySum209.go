func minSubArrayLen(target int, nums[]int) int{
  minLenght := len(nums) +1
  currentSum := 0
  left := 0
  for right := 0; right < len(nums); right++{
    currentSum += nums[right]
    for currentSum >= target{
      if right - left +1 < minLenght {
        minLenght = right - left +1
      }
      currentSum -= nums[left]
      left++
    }
  }
  if minLenght == len(nums){
    return 0
  }
  return minLenght
}
