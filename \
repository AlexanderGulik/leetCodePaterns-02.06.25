func sortList(head *ListNode) *ListNode{
  if head == nil || head.Next == nil{
    return head
  }
  
  mid := findMid(head)
  left := head
  right := mid.Next
  mid.Next = nil
  left = sortList(left)
  right = sortList(right)
  return merge(left, right)
}

func findMid(head *ListNode) *ListNode{
  slow, fast := head, head
  for fast.Next != nil && fast.Next.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
  }
  return slow
}

func merge(l1, l2 *ListNode) *ListNode{
  dummy := &ListNode{}
  current := dummy
  for l1 != nil && l2 != nil {
    if l1.Val < l2.Val{
      current.Next = l1
      l1 = l1.Next
    } else {
      current.Next = l2
      l2 = l2.Next
    }
    current = current.Next
  }
  if l1 != nil {
    current.Next = l1
  } else {
    current.Next = l2
  }
  return dummy.Next
}
