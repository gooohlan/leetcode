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

//      fbaonsiaxpruyg mzspoittzlhuhrys stumtch qgsgoptlmxko ydemhbfh fzufzzvjn o ywvtywdsmirgfwmiapf wynfwbi azayqvzvrkjwgxydozdv lpioshhwxqgnqsqdj b aqxvaycvb hkejxyxy ggz"
//      fbaonsiaxpruyg mzspoittzlhuhrys stumtch qgsgoptlmxko ydemhbfh fzufzzvjn o ywvtywdsmirgfwmiapf wynfwbi azayqvzvrkjwgxydozdv lpioshhwxqgnqsqdj bs aqxvaycvb hkejxyxy ggz


func Test648(t *testing.T) {
	dictionary := []string{
		"bsmlto",
	}
	fmt.Println(replaceWords(dictionary, "bs"))
}