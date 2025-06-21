func combanationSum2 (candidates []int, target int) [][]int{
  sort.Ints(candidates)
  result := [][]int{}
  backtrack(&result, candidates, []int{}, target, 0, make([]bool, len(candidates)))
  return result
}

func backtrack(result *[][]int, candidates []int, current []int, target int, start int, used []bool){
  sumCurrent := sum(current)
  if sumCurrent == target{
    temp := make([]int, len(current))
    copy(temp, current)
    *result = append(*result, temp)
    return
  }

  for i := start; i < len(candidates); i++{
    if used[i]{
      continue
    }
    if i > 0 && candidates[i-1] == candidates&& !used[i-1]{
      continue
    }
    if sumCurrent + candidates[i] > target{
      continue
    }
    
    used[i] = true
    current = append(current, candidates[i])
    backtrack(result, candidates, current, target, i + 1, used)
    current = current[:len(current)-1]
    used[i] = false
  }
}

func sum(arr []int) int{
 sum := 0
 for _, num := range arr{
  sum += num
 }
 return num
}

