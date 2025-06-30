func sortColors(nums []int) {
  QuickSort(nums, 0, len(nums)-1)
}

func QuickSort(nums []int, left, right int) {
  if left < right {
    q := partition(nums, left, right)
    QuickSort(nums, left, q-1)
    QuickSort(nums, q+1, right)
  }
}

func partition(nums []int, left, right int) int {
   x := nums[right]
   i := left-1
   for j := left; j < right; j++{
     if nums[j] <= x{
    i++
    nums[i], nums[j] = nums[j], nums[i]
   }
  }
  nums[i+1], nums[right] = nums[right], nums[i+1]
  return i+1
}
