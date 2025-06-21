func combinationSum3(k int, n int) [][]int{
  result := [][]int{}
  backtrack(&result, []int{}, 1, k, n)
  return result
}

func backtrack(result *[][]int, current []int, start, k, n int){
  if len(current) == k && n == 0{
    temp := make([]int, len(current))
    copy(temp, current)
    *result = append(*result, temp)
    return
  }

  for i := start; i <=9; i++{
    if n < i{
      continue
    }
    if len(current) +(9 - i + 1) < k {
      continue 
    }
    current = append(current, i)
    backtrack(result, current, i+1, k, n-i)
    current = current[:len(current) - 1]
  }
}
