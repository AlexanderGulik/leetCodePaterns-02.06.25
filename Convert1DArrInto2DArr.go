func construct2DArray(original[], n, m int) [][]int{
  if len(original) != n*m{
    return [][]int{}
  }
  result := make([][]int, m)
  for i := 0; i < m; i++{
    result[i] = make([]int, n)
    for j := 0;j < n; j++{
      result[i][j] = original[i*n+j]
    }
  }
  return result
}
