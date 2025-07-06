func smallestRange(nums [][]int) []int{
  type Pair struct {
    num, index int
  }
  var all []Pair
  for i, group := range nums {
    for _, num := range group {
      all = append(all, Pair{num, i})
    }
  }
  sort.Slice(all, func(i, j int) bool {
    return all[i].num < all[j].num
  })
  count := make(map[int]int)
  left, right := 0,0
  totalLists := len(nums)
  uniqueListsInWindow := 0
  minRange := int(^uint(0)>>1)
  result := []int{0,0}
  for right < len(all) {
    count[all[right].index]++
    if count[all[right].index] == 1 {
      uniqueListsInWindow++
    }
    for uniqueListsInWindow == totalLists {
      currentRange := all[right].num - all[left].num
      if currentRange < minRange {
        minRange = currentRange
        result = []{all[left].num, all[right].num}
      }
      count[all[left].index]--
      if count[all[left].index]==0 {
        uniqueListsInWindow--
      }
      left++
    }
    right++
  }
  return result
}

