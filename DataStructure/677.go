package DataStructure

// 直接引用1804的insert和CountWordsStartingWith即可
type MapSum struct {
    tireNode *Trie2
}


func Constructor677() MapSum {
    return MapSum{
        tireNode: &Trie2{},
    }
}


func (this *MapSum) Insert(key string, val int)  {
    this.tireNode.insert(key, val)
}


func (this *MapSum) Sum(prefix string) int {
    return this.tireNode.CountWordsStartingWith(prefix)
}
