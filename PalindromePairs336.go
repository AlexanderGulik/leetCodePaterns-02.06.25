func palindromePairs(words []string) [][]int {
    result := [][]int{}
    var isPalindromic func(string) bool
    isPalindromic = func(str string) bool {
        if len(str) < 2 {return true}
        left, right := 0, len(str) - 1
        for left < right {
            if str[left] != str[right] {
                return false
            }
            left++
            right--
        }
        return true
    }
    for i := 0; i<len(words); i++{
        for j :=0; j <len(words); j++{
            if i != j && isPalindromic(words[i] + words[j]){
                result = append(result, []int{i, j})
            }
        }
    }
    return result
}
