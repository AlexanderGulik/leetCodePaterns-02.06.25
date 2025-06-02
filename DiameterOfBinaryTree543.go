func diameterOfBinaryTree (root *TreeNode) int {
  var diameter int
  depth func (*TreeNode) int
  depth = func(node *TreeNode) int {
    if node == nil {
      retun 0
    }
    left := depth(node.Left)
    right := depth(node.Right)
    diameter = max(diameter, right+left)
    
    return 1 + max(left, right) 

  }
  depth(root)
  retun diameter

}

func max(a,b int) int {
  if a < b{
  return b
}
return a
  
}
