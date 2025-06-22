func detectCycle(head *ListNode) *ListNode {
  if head == nil || head.Next == nil{
    return nil
  }
  slow, fast := head, head
  for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
    if slow == fast{
      break
    }
  }
  if slow != fast{
    return nil
  }

  slow = head
  for slow != fast{
    fast = fast.Next
    slow = slow.Next
  }
  return slow
}
