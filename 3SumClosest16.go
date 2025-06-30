func threeSumClosest(nums []int, target int) int {
  result := nums[0] + nums[1] + nums[2]
  for i := 0;i < len(nums); i++{
    for j := i+1; j< len(nums); j++{
      for k := j+1; k<len(nums); k++{
        sum := nums[i] + nums[j] + nums[k]
        if sum == target{
          return sum
        }
        if (abs(sum-target) < abs(result-target){
          result = sum
        }
      }
    }
  }
  return result
}

func abs(x int) int{
   if x < 0{
      return -x
   }
   return x
}

//2

func threeSumClosest(nums []int, target int) int {
  sort.Ints(nums)
  closestSum := nums[0] + nums[1] + nums[2]

  for i := 0; i < len(nums)-2; i++{
    left, right := i+1, len(nums)-1
    for left < right {
      sum := nums[i] + nums[left] + nums[right]
      if sum == target {
        return sum
      }
      if abs(sum-target) < abs(closestSum-target){
        closestSum = sum
      }

      if sum < target {
        left++
      } else {
        right--
      }
    }
  }
  return closestSum
}
