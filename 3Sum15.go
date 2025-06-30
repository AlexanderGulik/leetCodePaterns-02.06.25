func threeSum(nums []int) [][]int {
  sort.Ints(nums)
  result := [][]int{}
  for i := 0; i < len(nums) - 2; i++{
    if i > 0 &&  nums[i-1] == nums[i] {
      continue
    }
    j, k := i+1, len(n)-1
    for j < k {
      sum := nums[i] + nums[j] + nums[k]
      if sum == 0 {
        result = append(result, []int{nums[i], nums[j], nums[k]})
        for j < k && nums[j] == nums[j+1]{
          j++
        }
        for j < k && nums[k] == nums[k-1] {
          k--
        }
        j++
        k--
      } else if sum < 0 {
        j++
      } else {
        k--
      }
    }
  }
  return result
}
