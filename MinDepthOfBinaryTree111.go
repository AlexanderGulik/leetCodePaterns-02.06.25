func minDepth(root *TreeNode) int {
  if root == nil{
    return 0
  }

  if root.Left == nil && root.Right == nil{
    return 1
  }

  if root.Left == nil {
      return minDepth(root.Right)+1
  }

  if root.Right == nil{
      return minDepth(root.Left)+1
  }

  DepthLeft := minDepth(root.Left)
  DepthRight := minDepth(root.Right)

  if DepthLeft < DepthRight {
    return DepthLeft + 1
  }
  return DepthRight + 1



}
