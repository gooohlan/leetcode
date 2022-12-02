package DataStructure

// 此题是208的拓展,在原有的基础上添加一些属性和方法即可
type Trie2 struct {
    children [26]*Trie2
    Val int
    isEnd bool // 标识节点是否结束
}

func Constructor1804() Trie2 {
    return Trie2{}
}

func (t *Trie2) Insert(word string)  {
    node := t.SearchPrefix(word)
    if node != nil && node.isEnd {
        node.Val ++
    }else{
        t.insert(word, 1)
    }
}

func (t *Trie2) insert(word string, value int){
    for _, ch := range word {
        ch -= 'a'
        if t.children[ch] == nil {
            t.children[ch] = &Trie2{}
        }
        t = t.children[ch]
    }
    t.Val = value
    t.isEnd = true
}

func (t *Trie2) CountWordsEqualTo(word string) int {
    node := t.SearchPrefix(word)
    if node == nil {
        return 0
    }
    return node.Val
}

// 找出所有前缀为word的节点,累加value
func (t *Trie2) CountWordsStartingWith(prefix string) int {
    root := t.SearchPrefix(prefix) // 获取前缀节点
    if root == nil {
        return 0
    }
    
    var (
        count int
        dfs func(trie2 *Trie2)
    )
    dfs = func(node *Trie2){
        if node == nil {
            return
        }
        
        count += node.Val
        for i := 0; i < 26; i++ { // 遍历子节点
            dfs(node.children[i])
        }
    }
    
    dfs(root)
    return count
}


func (t *Trie2) Search(word string) bool {
    node := t.SearchPrefix(word)
    return node != nil && node.isEnd
}


func (t *Trie2) StartsWith(prefix string) bool {
    node := t.SearchPrefix(prefix)
    return node != nil
}

// SearchPrefix 查找节点是否存在
func (t *Trie2) SearchPrefix(prefix string) *Trie2 {
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

func (t *Trie2) Erase(word string) {
    node := t.SearchPrefix(word)
    if node == nil || !node.isEnd {
        return
    }
    
    if node.Val == 1 {
        t.Remove(word)
    }else{
        node.Val--
    }
}

func (t *Trie2) Remove(word string) {
    
    var remove func (*Trie2, int) *Trie2
    remove = func(node *Trie2, index int) *Trie2 {
        if node == nil {
            return nil
        }
        if index == len(word) {
            node.Val = 0
            node.isEnd = false
        }else{
            ch := word[index] - 'a'
            node.children[ch] = remove(node.children[ch], index+1)
        }
        
        // 后序位置,处理递归路上节点
        if node.isEnd {
            return node
        }
    
        for i := 0; i < 26; i++ {
            if node.children[i] != nil { // 存在子节点,不会被清除
                return node
            }
        }
        
        // 既没有存储值,也没有子节点,清理掉
        return nil
    }
    t = remove(t, 0)
}
