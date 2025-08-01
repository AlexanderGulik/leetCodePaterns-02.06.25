func oddEvenList(head *ListNode) *ListNode{
  if head == nil || head.Next == nil{
    return head
  }
  odd := head
  even := head.Next
  HeadEven := even
  for even != nil && even.Next != nil{
    odd.Next = even.Next
    odd = odd.Next
    even.Next = odd.Next
    even = even.Next
  }
  odd.Next = HeadEven
  return head
}
