package DataStructure

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	constructor := Constructor2(3)
	constructor.Put(1, 1)
	constructor.Put(2, 2)
	constructor.Put(2, 2222)
	constructor.Put(3, 3)
	constructor.Put(4, 4)
	constructor.Get(2)
	constructor.Put(5, 5)
}

func TestLFU(t *testing.T) {
	constructor := Constructor460(0)
	constructor.Put(0, 0)
	constructor.Put(2, 2)
	constructor.Put(1, 1111)
	constructor.Get(2)
	constructor.Put(3, 3)
}

func Test208(t *testing.T) {
	trie  := Constructor208()
	trie.Insert("apple")
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.ShortestPrefixOf("adpbbccs"))
}

func Test648(t *testing.T) {
	dictionary := []string{
		"a", "b", "c",
	}
	fmt.Println(replaceWords(dictionary, "aadsfasf absbs bbab cadsfafs"))
	fmt.Println(replaceWords2(dictionary, "aadsfasf absbs bbab cadsfafs"))
}

func Test211(t *testing.T) {
	wordDictionary := Constructor211()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	fmt.Println(wordDictionary.Search("pad"))
	fmt.Println(wordDictionary.Search("bad"))
	fmt.Println(wordDictionary.Search(".ad"))
	fmt.Println(wordDictionary.Search("b.."))
}