func swapPairs(head *ListNode) *ListNode{
  dummy := &ListNode{Next: head}
  prev := dummy
  for prev.Next != nil && prev.Next.Next != nil{
    first := prev.Next
    second := first.Next
    nextNode := second.Next
    prev.Next = second
    second.Next = first
    first.Next = nextNode

    prev =  first
  }
  retun dummy.Next
}
