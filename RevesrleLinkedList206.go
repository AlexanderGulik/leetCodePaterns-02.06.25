package main

func reverseList(head *ListNode) *ListNode{
  var prev *ListNode
  current := head
  for current != nil{
  next:= current.Next
  current.Next = prev
  prev = current
  current = next
  }
  return prev

}
