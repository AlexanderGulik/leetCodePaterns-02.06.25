func pacificAtlantic(heights [][]int) [][]int {
  rows, cols := len(heights), len(heights[0])
  pacific := make([][]bool, rows)
  atlantic := make([][]bool, rows)

  for i := range cols {
    pacific[i] = make([]bool, cols)
    atlantic[i] = make([]bool, cols)
  }

  for i := 0; i < rows; i++{
    dfs(heights, pacific, i, 0, -1)
  }
  
  for j := 0; j < cols; j++{
    dfs(heights, pacific, 0, j, -1)
  }
  
  for i := 0; i< rows; i++{
    dfs(heights, atlantic, i, cols-1, -1)
  }

  for j := 0; j< cols; j++{
    dfs(heights, atlantic, rows-1, j, -1)
  }
  result := [][]int{}
  for i := 0; i < rows; i++{
    for j := 0; j < cols; j++{
      if atlantic[i][j] && pacific[i][j]{
        result = append(result, []int{i,j})
      }
    }
  }
  return result


}

func dfs(heights [][]int, visited [][]bool, i, j, prevHeight int){
    if i < 0 || i >= len(heights) || j < 0 || j >= len(heights[0]) || visited[i][j] || heights[i][j] < prevHeight{
        return
    }
    visited[i][j] = true
    currentHeight := heights[i][j]
    dfs(heights, visited, i+1, j, currentHeight)
    dfs(heights, visited, i-1, j, currentHeight)
    dfs(heights, visited, i, j+1, currentHeight)
    dfs(heights, visited, i, j-1, currentHeight)
}
