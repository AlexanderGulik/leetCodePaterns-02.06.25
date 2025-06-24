func checkInclusion(s1 string, s2 string) bool {
  if len(s1)> len(s2){
    return false
  }
  countS1 := [26]int{}
  countS2 := [26]int{}
  for i :=0; i < len(s1); i++{
    countS1[s1[i]-'a']++
    countS2[s2[i]-'a']++
  }
  matches := 0
  for i := 0; i < 26; i++{
    if countS2[i] == countS1[i]{
      matches++
    }
  }
  left := 0
  for right := len(s1); right < len(s2); right++{
    if matches == 26{
      return true
    }
    index := s2[right] - 'a'
    countS2[index]++
    if countS1[index] == countS2[index]{
      matches++
    } else if countS1[index] + 1 == countS2[index]{
      matches--
    }

    index = s2[left] - 'a'
    countS2[index]--
    if countS1[index] == countS2[index] {
      matches++
    } else if countS1[index] -1 == countS2[index] {
      matches--
    }
    left++
  }
  return 26 == matches
}
