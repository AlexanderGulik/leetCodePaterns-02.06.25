func letterCombinations(digits string) []string{
  if len(digits) == 0{
    return []string{}
  }
  result := []string{}
  letterMap := map[byte]string{
    '2' : "abc",
    '3' : "def",
    '4' : "ghi",
    '5' : "jkl",
    '6' : "mno",
    '7' : "pqrs",
    '8' : "tuv",
    '9' : "wxyz",
  }
  backtrack(&result, "", digits, letterMap, 0)
  return result
}

func backtrack(result *[]string, current string, digits string, letterMap map[byte]string, index int){
    if index == len(digits){
      *result = append(*result, current)
      return
    }
    letters := letterMap[digits[index]]
    for i := 0; i < len(letters); i++{
      backtrack(result, current + string(letters[i]), digits, letterMap, index+1)
    }
}
