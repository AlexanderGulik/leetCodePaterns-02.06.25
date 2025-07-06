func maxSlidingWindow(nums []int, k int) []int {
  if len(nums) < k || 0 == k {
    return make([]int,0)
  } else if 1 == k {
    return nums
  }
  var (
    result = make([]int, len(nums)-k+1)
    dq = &deque{}
  )
  for i := range nums {
    if false == dq.empty() && (i-k==dq.getFirst()){
      dq.popFirst()
    }
    for false == dq.empty() && nums[dq.getLast()] < nums[i] {
      dq.popLast()
    }
    dq.push(i)
    if i >= k-1{
      result[i-k+1] = nums[dq.getFirst()]
    }
  }
  return result
}

type deque struct {
  indexes []int
}

func (d *deque) push(i int){
  d.indexes = append(d.indexes, i)
}
func (d *deque) getFirst() int {
  return d.indexes[0]
}
func (d *deque) popFirst(){
  d.indexes = d.indexes[1:]
}
func (d *deque) getLast(){
  return d.indexes[len(d.indexes)-1]
}
func (d *deque)popLast(){
  d.indexes = d.indexes[:len(d.indexes)-1]
}
func (d *deque) empty() bool {
  return 0 == len(d.indexes)
}
