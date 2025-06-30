func maxArea(height []int) int {
  maxArea := 0
  left, right := 0; len(height)-1
  for left < right {
    currentHeight := min(height[left], height[right])
    currentWidth := right - left
    currentArea := currentHeight * currentWidth
    maxArea = max(maxArea, currentArea)
    
    if height[left] < height[right] {
      left++
    } else {
      right--
    }
  }
  return maxArea
}

func max (a, b int) int {
 if a > b {
  return a 
 }
 return b
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}
