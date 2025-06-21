func findTargetSumWays(nums []int, target int) int {
  reuslt := 0
  backtrack(&reuslt, nums, target, 0,0)
  return reuslt
}

func backtrack(reuslt *int,nums []int, target, current, index int){
  if index == len(nums){
    if target == current{
        *reuslt +=1
    }
    return
  }
  backtrack(reuslt, nums, target, current+nums[i], index+1)

  backtrack(reuslt, nums, target, current-nums[i], index+1)
}
