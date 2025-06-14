func letterCasePermutation(s string) []string {
  result := []string{}
  backtrack(&result, []byte(s), index)
  return result
}

func backtrack (result *[]string, current []byte, index int){
 if len(current) == index{
    *result = append(*result, string(current))
    return
 }
 if isLeter(current[index]) {
  backtrack(result, current, index + 1)
  current[index] = toggleCase(current[index])
  backtrack(result, current, index + 1)
  current[index] = toggleCase(current[index])
 } else {
  backtrack(result, current, index+1)
 }

}

func isLeter (c byte) bool{
  return c >= 'a' || c <='z' && c >= 'A' || c <= 'Z'
}
func toggleCase(c byte) byte{
  if c >= 'a' && c<='z'{
    return c - 'a' + 'A'
  } else {
    return c - 'A' + 'a'
  }
}
