func averageOfLevels(root *TreeNode) []float64{
  if root == nil {
    return []float64
  }

  var result []float64
  queue :=[]*TreeNode{root}
    for len(queue)>0{
      levelSize := len(queue)
      levelSum := 0.0
        for i:=0; levelSize>i; i++{
          node := queue[0]
          queue = queue[1:]
          levelSum += float64(node.Val)

          if node.Left != nil {
            queue = append(queue, node.Left)
          }
          if node.Right != nil{
            queue = append(queue, node,Right)
          }
        }
        average := levelSum / float64(levelSize)
        result = append(result, average)

    }
    return result
}
