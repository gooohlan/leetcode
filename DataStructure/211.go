package DataStructure


type WordDictionary struct {
    trieRoot *Trie
}


func Constructor211() WordDictionary {
    return WordDictionary{&Trie{}}
}


func (w *WordDictionary) AddWord(word string)  {
    w.trieRoot.Insert(word)
}


func (t *Trie) HasKeyWithPattern(pattern string, index int) bool {
    if t == nil {
        return false
    }
    if len(pattern) == index {
        return t.isEnd
    }
    
    ch := pattern[index]
    if ch != '.' {
        child := t.children[ch-'a']
        return child.HasKeyWithPattern(pattern, index+1)
    }
    
    for _, child := range t.children { // 遍历所有字母,只要有一个与之匹配就返回
         if child.HasKeyWithPattern(pattern, index+1) {
             return true
         }
    }
    return false
}

func (w *WordDictionary) Search(word string) bool {
    return w.trieRoot.HasKeyWithPattern(word, 0)
}
