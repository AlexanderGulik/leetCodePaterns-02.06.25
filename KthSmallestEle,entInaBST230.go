func kthSmallest(root *TreeNode, k int) int {
  stack := []*TreeNode{}
  current := root
  count := 0
  for current != nil || len(stack) > 0 {
    for current != nil {
      stack = append(stack, current)
      current = current.Left
    }
    current = stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    count++
     if count == k {
        return current.Val
     }
     current = current.Right
  }
  return -1
}
