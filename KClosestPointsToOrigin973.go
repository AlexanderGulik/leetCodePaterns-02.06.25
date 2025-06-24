func KClosest(points[][]int, k int) [][]int {
  type pointWithDist struct {
    dist int
    index int
  }
  var distances []pointWithDist
  for i, p := range points {
    distances = append(distances, pointWithDist{ 
      dist: p[0]*p[0] + p[1]*p[1],
      index: i,
  })}
  sort.Slice(distances, func(i, j int) bool{
    return distances[i].dist < distances[j].dist
  })
  
  result := make([][]int, k)
  for i := 0; i < k; i++{
    result[i] = points[distances[i].index]
  }
  return result

}
