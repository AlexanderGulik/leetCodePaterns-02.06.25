func isValidBST(root *TreeNode) bool {
  stack := []*TreeNode{}
  var prev *int = nil 
  for root != nil || len(stack) > 0 {
    for root != nil {
     stack = append(stack, root)
     root = root.Left
   }
   root = stack[len(stack)-1]
   stack = stack[:len(stack)-1]
   if prev != nil && root.Val <= *prev {
    return false
   }
   prev =&root.Val
   root =root.Right
  }
  return true
}
