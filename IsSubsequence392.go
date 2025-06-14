func isSubsequnce(s string, t string) bool{
  i, j := 0,0
  for i< len(s) && j< len(t){
    if s[i] == t[j]{
      i++
    }
    j++
  }
  return len(s) == i

}
