func findDuplicate(nums []int) int {
  seem := make(map[int]bool)
  for _, num := range nums {
    if seem[num] {
        return num
    }
    seem[num] = true
  }
  return -1
}

func findDuplicate2(nums []int) int {
  slow, fast := nums[0], nums[0]
  for {
    slow = nums[slow]
    fast = nums[nums[fast]]
    if slow == fast {
      break
    }
  }

  slow = nums[0]
  for slow != fast {
    slow = nums[slow]
    fast = nums[fast]
  }
  return slow
}
