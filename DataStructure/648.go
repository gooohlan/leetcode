package DataStructure

import "strings"

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
// 使用前缀树存储节点, 基于208题定义的前缀树添加一个最短前缀函数 ShortestPrefixOf
func replaceWords(dictionary []string, sentence string) string {
    root := Constructor208()
    for _, s := range dictionary {
        root.Insert(s)
    }
    
    words := strings.Split(sentence, " ")
    for i, word := range words {
        prefix := root.ShortestPrefixOf(word)
        if prefix != "" {
            words[i] = prefix
        }
    }
    return strings.Join(words, " ")
}

// 前缀树简化版
func replaceWords2(dictionary []string, sentence string) string {
    type trie map[rune]trie // 递归map替换前缀树
    root := trie{}
    for _, s := range dictionary {
        cur := root
        for _, c := range s {
            if cur[c] == nil {
                cur[c] = trie{}
            }
            cur = cur[c]
        }
        cur['#'] = trie{} // 结束标识
    }
    
    words := strings.Split(sentence, " ")
    for i, word := range words {
        cur := root
        for j, c := range word {
            if cur['#'] != nil { // 找到结束标记,表示找到最短前缀,替换当前值
                words[i] = word[:j]
                break
            }
            
            if cur[c] == nil { // 不存在前缀
                break
            }
            cur = cur[c]
        }
    }
    return strings.Join(words, " ")
}