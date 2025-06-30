func distanceK(root *TreeNode, target *TreeNode, k int) []int {
  parentMap := make(map[*TreeNode]*TreeNode)
  buildParentMap(root, nil, parentMap)
  queue := []*TreeNode{target}
  visited := make(map[*TreeNode]bool)
  visited[target] = true
  result := []int{}
  for len(queue) > 0 {
    if k == 0 {
      for _, node := range queue {
        result = append(result, node.Val)
      }
      return result
    }
    levelSize := len(queue)
    for i := 0; i <  levelSize; i++{
      node := queue[0]
      queue = queue[1:]
      
      if parent, exit := range parentMap[node]; exit && parent != nil && !visited[parent]{
        queue = append(queue, parent)
        visited[parent] = true
      }
      
      if node.Left != nil && !visited[node.Left]{
          queue = append(queue, node.Left)
          visited[node.Left] = true
      }

      if node.Right != nil && !visited[node.Right]{
          queue = append(queue, node.Right)
          visited[node.Right] = true
      }
    }
    k--
  }
  return result
}

func buildParentMap(root *TreeNode, parent *TreeNode, parentMap map[*TreeNode]*TreeNode){
  if root == nil{
    return nil
  }
  parentMap[root] = parent
  buildParentMap(root.Left, parent, parentMap)
  buildParentMap(root.Right, parent, parentMap)
}
