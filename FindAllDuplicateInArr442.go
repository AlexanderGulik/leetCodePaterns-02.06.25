func findDuplicates (nums []int) []int{
  result := []int{}
  seen := make(map[int]bool)
  for _, num := range nums{
    if seen[num]{
        result = append(result, num)
    }
    seen[num] = true
  }
  return result
}

func findDuplicates2(nums []int) []int {
  var result []int
  for _, num := range nums {
    idx := abs(num) -1

    if num[idx] < 0 {
      result = append(result, idx+1)
    } else {
      nums[idx] = -nums[idx]
    }
  }
  return result
}

func abs (n int) int{
  if n < 0{
    return -n
  }
  return n

}
