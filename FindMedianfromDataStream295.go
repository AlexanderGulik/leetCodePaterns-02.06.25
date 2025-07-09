type MaxHeap []int
func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j]}
func (h MaxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MaxHeap) Push(x interface{}) {*h = append(*h, x.(int))}
func (h *MaxHeap) Pop() interface{} {x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x}

type MinHeap []int
func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j]}
func (h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MinHeap) Push(x interface{}) {*h = append(*h, x.(int))}
func (h *MinHeap) Pop() interface{} {x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x}

type MedianFinder struct {
    low *MaxHeap
    high *MinHeap
}


func Constructor() MedianFinder {
    l, h := &MaxHeap{}, &MinHeap{}
    heap.Init(l)
    heap.Init(h)
    return MedianFinder{l,h}
}


func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.low, num)
    heap.Push(this.high, heap.Pop(this.low))
    if this.low.Len() < this.high.Len(){
        heap.Push(this.low, heap.Pop(this.high))
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.low.Len() > this.high.Len(){
        return float64((*this.low)[0])
    }
    return (float64((*this.low)[0])+ float64((*this.high)[0]))/ 2.0
}

