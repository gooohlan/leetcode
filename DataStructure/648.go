package DataStructure

import "strings"
// 使用前缀树存储节点
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