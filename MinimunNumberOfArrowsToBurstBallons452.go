func findMinArrowShots(points [][]int) int {
  if len(points) == 0 {
    return 0
  }

  sort.Slice(points, func(i, j int) bool {
    return points[i][1] < points[j][1]
  }
  arrow := 1
  shot := points[0][1]
  for i := 1; i < len(points); i++{
    if shot < points[i][0]{
      arrow++
      shot = points[i][1]
    }
  }
  return arrow
}
