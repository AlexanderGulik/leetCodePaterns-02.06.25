func rotateRight(head *ListNode, k int) *ListNode{
  if head == nil || head.Next == nil || k == 0{
    return head
  }
  tail := head
  Lenght := 1
  for tail.Next != nil{
    tail = tail.Next
    Lenght++
  }
  k = k % Lenght
  if k == 0 {
    return head
  }

  newTail := head 
  for i := 0; i < Lenght-k-1; i++{
    newTail = newTail.Next
  }
  headTail := newTail.Next
  newTail.Next = nil
  tail.Next = head
  return headTail

}
