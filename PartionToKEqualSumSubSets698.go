func canPartitionSubsets(nums []int, k int) bool{
  totalSum := 0
  for _, num := range nums{
    totalSum += num
  }
  if totalSum % k != 0{
    return false
  }
  target := totalSum / k
  sort.Sort(sort.Reverse(sort.IntSlice(nums)))
  used := make([]bool, len(nums))
  return backtrack(nums, used, k, 0, 0, target)

}

func backtrack(nums []int, used []bool, k, start, currentSum, target int) bool{
  if k == 1{
    return true
  }
  if currentSum == target {
    return backtrack(nums, used, k-1, 0,0,target)
  }
  for i := start; i < len(nums); i++{
    if !used[i] && nums[i] + currentSum <= target{
      used[i] = true
      if backtrack(nums, used, k, i + 1, nums[i]+currentSum, target) {
        return true
      }
      used[i] = false
    }
  }
  return false
}

