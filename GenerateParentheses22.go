func generateParenthesis(n int) []string {
  result := []string{}
  backtrack(&result,"", 0,0,n)
  return result
}

func backtrack(result *[]string, current string, open, clos, max) {
  if len(current) == max*2{
    *result = append(*result, current)
    return
  }
  if open < max {
    backtrack(result, current+"(", open+1,clos,max)
  }
  if open > clos{
    backtrack(result, current+")", open, clos+1, max)
  }
}
