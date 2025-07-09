type WordFilter struct {
    hash map[string]int
}


func Constructor(words []string) WordFilter {
    hash := make(map[string]int, len(words)*100)
    for i, word := range words {
        for l := 1; l <=10 && l <= len(word); l++ {
            for r := 1; r <= 10 && r <= len(word);r++{
                hash[word[:l] + ":"+word[len(word)-r:]] = i
            }
        }
    }
    return WordFilter {
        hash: hash,
    }
}


func (this *WordFilter) F(pref string, suff string) int {
    if v, exists := this.hash[pref+":"+suff]; exists {
        return v
    }
    return -1
}


/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
