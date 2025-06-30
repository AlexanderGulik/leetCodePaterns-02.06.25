func pathSum(root *TreeNode, targetSum int) int {
  if root == nil {
    return 0
  }
  return dfs(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func dfs(root *TreeNode, targetSum int) int {
  if root == nil {
    return 0
  }
  count := 0
  if node.Val == targetSum {
    count++
  }
  count += dfs(root.Left, targetSum-root.Val)
  count += dfs(root.Right, targetSum-root.Val)
  return count
}

