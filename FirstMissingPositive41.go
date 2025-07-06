func firstMissingPositive(nums []int) int {
  used := make([]bool, len(nums)+1)
  for _, num := range nums {
    if num > 0 && num < len(used){
      used[num] = true
    }
  }

  for i := 1; i < len(used); i++{
    if !used[i] {
      return i
    }
  }
  return len(used

}
