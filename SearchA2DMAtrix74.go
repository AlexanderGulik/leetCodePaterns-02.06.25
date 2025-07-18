func searchMatrix(matrix [][]int, target int) bool{
  rows := len(matrix)
  cols := len(matrix[0])
  left, right := 0, rows *cols-1
  for left <= right{
    mid := left +(right- left)/2
    midValue := matrix[mid/cols][mid%cols]
    if midValue == target{
      return true
    } else if midValue < target {
      left = mid +1
    } else {
      right = mid -1
    }
  }
  return false
}
