package main

func removeElements(head *ListNode, val int) *ListNode{
  dummy := &Listnode {Next: head}
  prev := dummy
  current := head
  for current != nil{
    if( current.Val = val){
      prev.Next = current.Next
    } else{
      prev = current
    }
    current = current
  }
  return dummy.Next


}
