func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
  if root == nil{
    return false
  }
  if isSameTree (root, subRoot){
      return true
  }

  return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)

}

func isSameTree (t *TreeNode, s *TreeNode) bool {
  if t == nil && s == nil {
    return true
  }
  if t== nil || s == nil {
      return false
  }

  return t.Val==s.Val && isSameTree(t.Left, s.Left) && isSameTree(t.Right, s.Right)
}
