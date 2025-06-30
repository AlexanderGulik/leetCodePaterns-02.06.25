func levelOrder(root *TreeNode) [][]int {
  if root == nil {
      return nil
  }
  queue := []*TreeNode{root}
  result := [][]int{}
  for len(queue) > 0 {
    levelSize := len(queue)
    currentVal := []int{}
    for i := 0; i < levelSize; i++{
      node := queue[0]
      queue = queue[1:]
      currentVal = append(currentVal, node.Val)
      if node.Left != nil{
        queue = append(queue, node.Left)
      }
      if node.Right != nil {
        queue = append(queue. node.Right)
      }
    }
    result = append(result, currentVal)
  }
  return result

}
