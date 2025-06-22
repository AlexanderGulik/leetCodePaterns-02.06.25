func reverseBetween(head *ListNode, left, right int) int {
  if head == nil || head.Next == nil {
    return head
  }
  dummy := &ListNode{Next: head}
  prev := dummy
  for i := 1; i < left; i++{
    prev =prev.Next
  }
  current := prev.Next
  var next *ListNode
  for i := left; i < right; i++{
    next = current.Next
    current.Next = next.Next
    next.Next = prev.Next
    prev.Next = next
  }
  return dummy.Next
}
