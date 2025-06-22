func reorderList(head *ListNode){
  if head != nil || head.Next != nil{
    return 
  }
  slow, fast := head, head
  for fast.Next != nil && fast.Next.Next != nil{
    slow = slow.Next
    fast = fast.Next.Next
  }
  headReverse := reverse(slow.Next)
  slow.Next = nil
  merge(head, headReverse)

}

func reverse(head *ListNode) *ListNode{
  var prev *ListNode
  current := head
  for current != nil {
    next := current.Next
    current.Next = prev
    prev = current
    current = next
  }
  return prev
}

func merge(l1, l2 *ListNode) {
  for l1 != nil && l2 != nil{
      nextL1 := l1.Next
      nextL2 := l2.Next

      l1.Next = l2
      l2.Next = nextL1
      l1 = nextL1
      l2 = nextL2
  }
}
