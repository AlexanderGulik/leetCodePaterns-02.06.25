func levelOrderBottom(root *TreeNode) [][]int{
  if root == nil {
    return [][]int{}
  }
  result := [][]int{}
  queue := []*TreeNode{root}
  for len(queue) > 0{
    levelSize := len(queue)
    currentVal := []int{}
    for i := 0; i < levelSize; i++{
      node := queue[0]
      queue = queue[1:]
      currentVal = append(currentVal, node.Val)
      if node.Left != nil {
        queue = append(queue, node.Left)
      }
      if node.Right != nil {
        queue = append(queue, node.Right)
      }
    }
    result = append(result, currentVal)
  }
  reverseArr(result)
  return result
}

func reverseArr(s [][]int) [][]int{
  for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1{
    s[i], s[j] = s[j], s[i]
  }
  return s
}
