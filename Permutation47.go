func permutation (nums []int) [][]int{
  result := [][]int{}
  sort.Ints(nums)
  backtrack(&result, nums, []int{}, make([]bool, len(nums)))
  return result
}

func backtrack (result *[][]int, nums []int, current []int, used []bool){
  if len(current) == len(nums){
    temp := make([]int, len(current))
    copy(temp, current)
    *result = append(*result, temp)
    return
  }
  for i := 0; i < len(nums); i++{
    if used[i]{
      continue 
    }
    if i > 0 && nums[i] == nums[i-1] && !used[i-1]{
      continue
    }
    used[i] = true
    current = append(current, nums[i])
    backtrack(result, nums, current, used)
    current = current[:len(current)-1]
    used[i] = false
  }
}
