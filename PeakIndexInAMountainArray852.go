func peakIndexInMountainArray(arr []int) int {
  for i := len(arr) - 1; i > 0; i--{
    if arr[i] > arr[i-1] {
      return i
    }
  }
  return 0
}// O(n)

func peakIndexInMountainArray (arr []int) int {
  left, right := 0, len(arr)-1
  for left < right {
    mid := left + (right-left) / 2
    if arr[mid] < arr[mid + 1] {
      left = mid + 1
    } else {
      right = mid
    }
  }
  return left
}
