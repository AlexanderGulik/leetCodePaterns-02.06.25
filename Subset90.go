func subsetWithDup(nums []int) [][]int{
  sort.Ints(nums)
  result := [][]int{}
  backtrack(&result, nums, []int{}, 0)
  return result
}

func backtrack(result *[][]int, nums []int, current []int, start int){
  temp := make([]int, len(current))
  copy(temp, current)
  *result = append(*result, temp)
  for i := start; i < len(nums); i++{ 
    if i > start && nums[i] == nums[i-1]{
      continue
    }
    current = append(current, nums[i])
    backtrack(result, nums, current, i+1)
    current = current[len(current)-1]
  }

}
