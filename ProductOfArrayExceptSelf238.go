func productExceptSelf(nums []int) []int {
  n := len(nums)
  arr := make([]int, n)
  arr[0] = 1
  for i := 1; i < n; i++ {
    arr[i] = arr[i-1]*nums[i-1]
  }
  sufix := 1
  for i := n-1; i >= 0; i--{
    arr[i] *= sufix
    sufix *= nums[i]
  }
  return arr
}
