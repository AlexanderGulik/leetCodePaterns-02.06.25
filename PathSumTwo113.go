func pathSum(root *TreeNode, targetSum int) [][]int {
  result := [][]int{}
  dfs(&result, root, targetSum, []int{})
  return result
}

func dfs(result *[][]int, node *TreeNode, targetSum int, current []int{}) {
  if node == nil {
    return
  }

  current = append(current, node.Val)
  if node.Left == nil && node.Right == nil && node.Val == targetSum {
    pathCopy := make([]int, len(current))
    copy(pathCopy, current)
    *result = append(*result, pathCopy)
    return
  }
  dfs(result, node.Left, targetSum-node.Val, current)
  dfs(result, node.Right, targetSum-node.Val, current)
}
