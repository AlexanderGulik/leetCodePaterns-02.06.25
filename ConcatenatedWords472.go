func findAllConcatenatedWordsInADict(words []string) []string {
    result := make([]string, 0)
    wordMap := make(map[string]struct{},0)
    for _, word := range words {
        wordMap[word] = struct{}{}
    }
    for _, word := range words {
        if wordBreak(word, wordMap){
            result = append(result, word)
        }
    }
    return result
}

func wordBreak(s string, wordMap map[string]struct{}) bool {
    wb := make([]bool, len(s) + 1)
    wb[0] = true
    for i := 1; i < len(s) + 1; i++{
        for j := 0 ; j < i; j++ {
            toSearch := s[j:i]
            _, found := wordMap[toSearch]
            if (toSearch != s && found && wb[j]){
                wb[i] = true
                break
            }
        }
    }
    return wb[len(s)]
}
