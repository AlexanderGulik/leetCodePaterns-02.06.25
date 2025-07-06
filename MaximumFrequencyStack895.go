type FreqStack struct {
  s [][]int
  h map[int]int
}

func Constructor() FreqStack{
  return FreqStack{
    h: make(map[int]int),
  }
}

func (this *FreqStack) Push(val int) {
  this.h[val]++
  if len(this.s) < this.h[val] {
    this.s = append(this.s, []int{val})
    return
  }
  this.s[this.h[val]-1] = append(this.s[this.h[val]-1], val)
}

func (this *FreqStack) Pop() int {
  top := len(this.s)-1
  val := this.s[top][len(this.s[top])-1]
  this.s[top] = this.s[top][:len(this.s[top])-1]
  if len(this.s[top]) == 0 {
    this.s =  this.s[:top]
  }
  this.h[val]--
  return val
}
