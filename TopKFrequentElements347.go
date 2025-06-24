func topKFrequent(nums []int, k int) []int {
  freq := make(map[int]int)
  for _, num := range nums {
    freq[num]++
  }
  type kv struct {
    key int
    value int
  }
  var freqList []kv
  for key, value := range freq {
    freqList = append(freqList, kn{key, value})
  }
  sort.Slice(freqList, func (i, j int) bool{
    return freqList[i].value > freqList[j].value
  })
  result := make([]int, k)
  for i := 0; i<k;i++{
    result[i] = freqList[i].key
  }
  return result
}
