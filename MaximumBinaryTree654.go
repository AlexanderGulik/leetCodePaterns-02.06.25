func constructMaximumBinaryTree(nums []int) *TreeNode {
  if len(nums) == 0 {
    return nil
  }

  root := &TreeNode{
    Val: nums[indexMax],
    Left: constructMaximumBinaryTree(nums[:indexMax]),
    Right: constructMaximumBinaryTree(nums[indexMax+1:]),
  }

  return root
}

func findMaxIndex(arr []int) int {
  indexMax := 0
  for i := 1; i < len(arr); i++{
    if arr[i] > arr[indexMax] {
      indexMax = i
    }
  }
  return indexMax
}
