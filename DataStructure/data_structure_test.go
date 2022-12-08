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
	trie := Constructor208()
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

func Test1804(t *testing.T) {
	trie := Constructor1804()
	trie.Insert("abbce")
	trie.Insert("abbce")
	trie.Insert("abb")
	fmt.Println(trie.CountWordsEqualTo("abbce"))
	fmt.Println(trie.CountWordsStartingWith("abb"))
	trie.Erase("abbce")
	fmt.Println(trie.CountWordsEqualTo("abbce"))
	fmt.Println(trie.CountWordsStartingWith("abb"))
	trie.Erase("abb")
	fmt.Println(trie.CountWordsEqualTo("abbce"))
	fmt.Println(trie.CountWordsStartingWith("abb"))
	trie.Erase("abbce")
	fmt.Println(trie.CountWordsEqualTo("abbce"))
	fmt.Println(trie.CountWordsStartingWith("abb"))
}

func Test677(t *testing.T) {
	mapSum := Constructor677()
	mapSum.Insert("apple", 3)
	fmt.Println(mapSum.Sum("ap"))
	mapSum.Insert("app", 2)
	mapSum.Insert("apple", 2)
	fmt.Println(mapSum.Sum("ap"))
}

func Test295(t *testing.T) {
	this := Constructor295()
	this.AddNum(1)
	this.AddNum(2)
	this.AddNum(3)
	this.AddNum(4)
	this.AddNum(5)
	fmt.Println(this.FindMedian())
}

func Test496(t *testing.T) {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

func Test739(t *testing.T) {

}
