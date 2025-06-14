func permute(nums []int) [][]int {
  result := [][]int{}
  backtrack(&result, nums, []int{}, make([]bool, len(nums))
  return result
} 

func backtrack( result *[][]int, nums []int, current []int, used []bool){
  if len(current) == len(nums){
    temp := make([]int, len(current))
    copy(temp, current)
    *result = append(*result, temp)
    return
  }

  for i := 0; i <len(nums); i++{
    if !used[i]{
      used[i] = true
      current = append(current, nums[i])
      backtrack(result, nums, current, used)
      current = current[:len(current)-1]
      used[i] = false
    }
  }
}
