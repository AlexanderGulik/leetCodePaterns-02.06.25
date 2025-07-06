func countRangeSum(nums []int, lower, upper int) int {
  n := len(nums)
  prefix := make([]int, n+1) 
  for i := 0; i < n; i++{
    prefix[i+1] = prefix[i] + nums[i]
  }
  sortedPrefix := make([]int, len(prefix))
  copy(sortedPrefix, prefix)
  sort.Ints(sortedPrefix)
  ft := NewFenwickTree(len(sortedPrefix))
  result := 0
  for _, p := range prefix {
    lowerBound := p - upper
    upperBound := p - lower
    result += ft.Query(indexOf(sortedPrefix, upperBound+1)) - ft.Query(indexOf(sortedPrefix, lowerBound))
    ft.Update(indexOf(sortedPrefix, p)+1, 1)
  }
  return result
}

func FenwickTree struct {
  tree []int
}

func NewFenwickTree(n int) *FenwickTree {
  return &FenwickTree{
    tree: make([]int, n+1),
  }
}
func (ft *FenwickTree) Update(i, delta int) {
  for i < len(ft.tree) {
    ft.tree[i] ++ delta
    i += i & -i
  }
}

func (ft *FenwickTree) Query(i int) int {
  sum := 0
  for i > 0 {
    sum += ft.tree[i]
    i -= i & -i
  }
  return sum
}

func indexOf(sorted []int, value int) int {
  return sort.SearchInts(sorted, value)
}
