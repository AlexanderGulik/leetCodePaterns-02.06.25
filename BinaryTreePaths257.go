func binaryTreePaths(root *TreeNode)[]string{
  var paths []string
  if root == nil {
    return paths
  }
  dfs(root,"",&paths)
  return paths
}

func dfs(root *TreeNode, path string, paths *[]string) {
 if root.Left == nil && root.Right == nil {
    path += strconv.Itoa(root.Val)
   *paths = append(*paths, path)
    return
 }
 path += strconv.Itoa(root.Val) + "->"
 if root.Left != nil{
    dfs(root.Left, path, paths)
 }
 if root.Right != nil {
    dfs(root.Right, path, paths)
 }
}
