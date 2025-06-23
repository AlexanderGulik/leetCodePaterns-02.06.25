func merge(intervals [][]int) [][]int {
  sort.Slice(intervals, func(i, j int)bool {
    return intervals[i][0] < intervals[j][0]
  })
  merged := [][]int{intervals[0]}
  for i := 1; i < len(intervals); i++{
    last := merged[len(merged)-1]
    current := intervals[i]
    if last[1] >= current[0]{
      if last[1] < current[1]{
        last[1] = current[1]
      }
    } else {
      merged = append(merged, current)
    }
  }
  return merged
}
