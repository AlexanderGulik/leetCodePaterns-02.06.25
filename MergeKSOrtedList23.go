func mergeKLists(lists []*ListNode) *ListNode {
  if len(lists) == 0 {
    return nil
  }
  for len(lists) > 1 {
    merged := []*ListNode{}
    for i := 0; i < len(lists); i +=2{
      if i+1 < len(lists) {
        merged = append(merged, mergeTwo(lists[i], lists[i+1]))
      } else {
        merged = append(merged, lists[i])
      }
    }
    lists = merged
  }
  return lists[0]
}

func mergeTwo(l1, l2 *ListNode) *ListNode {
  dummy := &ListNode{}
  temp := dummy
  for l1 != nil && l2 != nil {
    if l1.Val <= l2.Val {
      temp.Next = l1
      l1 = l1.Next
    } else {
      temp.Next = l2
      l2 = l2.Next
    }
    temp = temp.Next
  }
  if l1 != nil {
    temp.Next = l1
  } else {
    temp.Next = l2
  }
  return dummy.Next
}
