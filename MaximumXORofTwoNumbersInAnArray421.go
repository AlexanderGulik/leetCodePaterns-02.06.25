type TrieNode struct {
  children [2]*TrieNode
}

func findMaximumXOR(nums []int) int {
  root := &TrieNode{}
  maxResult := 0

  for _, num := range nums {
    node := root
    for i := 31; i >= 0; i--{
      bit := (num >> uint(i)) & 1
      if node.children[bit] == nil {
        node.children[bit] = &TrieNode{}
      }
      node = node.children[bit]
    }
  }

  for _, num := range nums {
    node := root
    currentMax := 0
    for i := 31; i >= 0; i--{
      bit :=(num>> uint(i)) & 1
      toggledBit := 1-bit
      if node.children[toggledBit] != nil {
        currentMax += (1 << uint(i))
        node = node.children[toggledBit]
      } else {
        node = node.children[bit]
      }
    }
    if currentMax > maxResult {
      maxResult = currentMax
    }
  }
  return maxResult
}
