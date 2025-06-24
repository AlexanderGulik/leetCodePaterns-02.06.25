func findOrder(numCourses int, prerequisites [][]int) []int {
  graph := make(map[int][]int)
  inDegree := make([]int, numCourses)
  for _, p := range prerequisites {
    a, b := p[0], p[1]
    graph[b] = append(graph[b], a)
    inDegree[a]++
  }
  queue := []int{}
  for i := 0; i < numCourses; i++{
    if inDegree[i] == 0 {
      queue = append(queue, i)
    }
  }

  result := []int{}
  for len(queue) > 0 {
    course := queue[0]
    queue = queue[1:]
    result = append(result, course)
    for _, neighbor := range graph[course] {
      inDegree[neighbor]--
      if inDegree[neighbor] == 0 {
        queue = append(queue, neighbor)
      }
    }
  }
  if len(result) != numCourses{
    return []int{}
  }
  return result
}
