func twoSum(nums []int, target int) []int {
  var slice []int
  outerLoop:
  for i := 0; i < len(nums); i++{
    for j := 0; j < len(nums); j++{
      if ((nums[i]+nums[j]==target)&&i != j){
        slice = append(slice, i, j)
        break outerLoop
      }
    }
  }
  return slice
}

func twoSum(nums []int, target int) []int {
  hash := make(map[int]int)
  for i := 0; i< len(nums); i++{
    remain := target - nums[i]
    _, ok := hash[remain]
    if ok {
      return []int{hash[remain], i}
    }
    hash[nums[i]] = i
  }
  return []int{0,0}
}
