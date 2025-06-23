func insert(intervals [][]int, newInterval []int) [][]int {
  result := [][]int{}
  i := 0
  for i < len(intervals) && intervals[i][1] < newInterval[0] {
    result = append(result, intervals[i])
    i++
  }

  merged := []int{newInterval[0], newInterval[1]}
  for i < len(intervals) && intervals[i][0] <= newInterval[1] {
    merged[0] = min(merged[0], intervals[i][0])
    merged[1] = max(merged[1], intervals[i][1])
    i++
  }
  result = append(result, merged)
  for i < len(intervals[i]){
    result = append(result, intervals[i])
    i++
  }

  return result
}
func min (a,b int) int{
  if a < b{
    return a
  }
  return b
}
func max (a, b int) int {
  if a > b {
    return a
  }
  return b
}
