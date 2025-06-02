func hasPathSum(root *TreeNode, targetSum int) bool {
  if root == nil{
    return false
  }
  if root.Left == nil && root.Right == nil{
    return root.Val ==targetSum
  }
  return hasPathSum(root.Left, targetSum-root.Val)||(root.Right, targetSum-root.Val)
}
