type Trie struct {
  children[26]*Trie
  isWord bool
}

func Constructor() Trie {
  return Trie{}
}

func (this *Trie) Insert(word string) {
  node := this
  for _, ch := range word {
    idx := ch - 'a'
    if node.children[idx] == nil {
      node.children[idx] = &Trie{}
    }
    node = node.children[idx]
  }
  node.isWord = true  
}

func (this *Trie) Search(word string) bool {
  node := this.Find(word)
  return node != nil && node.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
  return this.Find(prefix) != nil
}

func (this *Trie) Find (s string) *Trie {
  node := this
  for _, ch := range s {
    idx := ch - 'a'
    if node.children[idx] == nil {
      return nil
    }
    node = node.children[idx]
  }
  return node
}
/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
