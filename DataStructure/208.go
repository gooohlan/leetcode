package DataStructure


type Trie struct {
    children [26]*Trie
    isEnd bool // 标识节点是否结束
}


func Constructor208() Trie {
    return Trie{}
}


func (t *Trie) Insert(word string)  {
    // node := t
    for _, ch := range word {
        ch -= 'a'
        if t.children[ch] == nil {
            t.children[ch] = &Trie{}
        }
        t = t.children[ch]
    }
    t.isEnd = true
}

func (t *Trie) Search(word string) bool {
    node := t.SearchPrefix(word)
    return node != nil && node.isEnd
}


func (t *Trie) StartsWith(prefix string) bool {
    node := t.SearchPrefix(prefix)
    return node != nil
}

// SearchPrefix 查找节点是否存在
func (t *Trie) SearchPrefix(prefix string) *Trie {
    // node := t
    for _, ch := range prefix {
        ch -= 'a'
        if t.children[ch] == nil {
            return nil
        }
        t = t.children[ch]
    }
    return t
}



func (t *Trie) ShortestPrefixOf(word string) string {
    for i, ch := range word {
        ch -= 'a'
        if t.children[ch] == nil {
            return ""
        }
        if t.children[ch].isEnd{
           return  word[:i+1]
        }
        t = t.children[ch]
    }
    return ""
}
