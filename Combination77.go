func combine(n int, k int) [][]int{
 result := [][]int{}
 backtrack(&result, []int{}, 1, n, k)
 return result
}

func backtrack(result *[][]int, current []int, start, n, k int){
  if k == len(current){
    temp := make([]int, len(current))
    copy(temp, current)
    *result = append(*result, temp)
    return
  }

  for i := start; i <= n; i++{
    current = append(current, i)
    backtrack(result, current, i+1, n, k)
    current = current[:len(current)-1]
  }
}
