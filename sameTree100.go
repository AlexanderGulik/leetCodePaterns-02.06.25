func isSameTree(p *TreeNode, q *TreeNode) bool{
  if q == nil && p == nil{
    return true
  }
  if q == nil || p == nil{
    return false
  }
  if q.Val != p.Val {
      return false
  }

  return isSameTree(p.Left, q.Left)&& isSameTree(p.Right, q.Right)
}
