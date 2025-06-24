func canFinish(numCourses int, prerequisites [][]int) bool {
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
  count := 0
  for len(queue) > 0 {
    node := queue[0]
    queue = queue[1:]
    count++

    for _, neighbor := range graph[node] {
      inDegree[neighbor]--
      if inDegree[neighbor] == 0 {
        queue = append(queue, neighbor)
      }
    }
  }
  return count = numCourses
}
