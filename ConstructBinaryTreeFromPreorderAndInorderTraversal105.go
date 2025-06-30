func buildTree(preorder []int, inorder[int]) *TreeNode {
  if len(preorder) == 0 || len(inorder) == 0 {
    return nil
  }
  rootVal := preorder[0]
  root := &TreeNode{Val: rootVal}
  var rootPoss int
  for i, val := range inorder {
    if val == rootVal {
      rootPoss = i
    }
  }
  leftInorder := inorder[:rootPoss]
  rightInorder := inorder[rootPoss+1:]
  
  leftPreorder := preorder[1: 1+len(leftInorder)]
  rightPreorder := preorder[1+len(leftInorder):]
  
  root.Left = buildTree(leftPreorder, leftInorder)
  root.Right = buildTree(rightPreorder, rightInorder)
  return root
}
