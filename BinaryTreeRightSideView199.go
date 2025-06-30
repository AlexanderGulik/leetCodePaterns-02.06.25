func rightSideView(root *TreeNode) []int {
  result := []int{}
  queue := []*TreeNode{}
  for len(queue) > 0 {
    levelSize := len(queue)
    result = append(result, queue[0].Val)
    for i := 0; i < levelSize; i++{
      node := queue[0]
      queue = queue[1:]
      if node.Right != nil {
        queue = append(queue, node.Right)
      }

      if node.Left != nil {
        queue = append(queue, node.Left)
      }
    }
  }
  return result
}
