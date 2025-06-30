func widthOfBinaryTree(root *TreeNode) int {
  if root == nil {
    return 0
  }
  maxWidth := 0
  queue := []struct{
    node *TreeNode
    index int
  } {{root, 1}}

  for len(queue) > 0 {
    levelSize := len(queue)
    left := queue[0].index
    right := left
    for i := 0; i< levelSize; i++{
      current := queue[0]
      queue = queue[1:]
      right = current.index

      if current.node.Left != nil {
        queue = append(queue, struct{
        node *TreeNode
        index int}{current.node.Left, 2*current.index})
      }
      if current.node.Right != nil {
        queue = append(queue, struct{
        node *TreeNode
        index int}{current.node.Right, 2*current.index+1})
      }
      currentWidth := right - left +1
      if currentWidth > maxWidth {
        maxWidth = currentWidth
      }

    }
  }
  return maxWidth
} 
