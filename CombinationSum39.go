func combinationSum(candidates []int, target[]int) [][]int{
  sort.Ints(candidates)
  result := [][]int{}
  backtrack(&result, candidates, []int{}, target, 0)
  return result
}

func backtrack(result *[][]int, candidates []int, current []int, target int, start int) {
 sumCurrent := sum(current)
 if sumCurrent == target{
    temp := make([]int, len(current))
    copy(temp, current)
    *result = append(*result, temp)
    return
 }

 for i := start; i < len(candidates); i++{
    if sumCurrent + candidates[i] > target{
      continue
    }
    current = append(current, candidates[i])
    backtrack(result, candidates, current, target, i)
    current = current[:len(current)-1]
 }
}

func sum(arr []int)int{
  sum := 0
  for _, num := range arr{
    sum += num
  }
  return sum
}
