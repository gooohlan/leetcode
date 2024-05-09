package main

import "strings"

// map处理判断是否在同一行
func findWords(words []string) []string {
	var (
		res []string
		keyboard map[byte]int
	)
	keyboard = map[byte]int{
		'q': 0, 'w': 0, 'e': 0, 'r': 0, 't': 0, 'y': 0, 'u': 0, 'i': 0, 'o': 0, 'p': 0,
		'a': 1, 's': 1, 'd': 1, 'f': 1, 'g': 1, 'h': 1, 'j': 1, 'k': 1, 'l': 1,
		'z': 2, 'x': 2, 'c': 2, 'v': 2, 'b': 2, 'n': 2, 'm': 2,
	}
	for _, word := range words {
		b := true
		lowercase := strings.ToLower(word)
		for i := 1; i < len(lowercase); i++ {
			if keyboard[lowercase[i-1]] != keyboard[lowercase[i]] {
				b = false
				break
			}
		}
		if b {
			res = append(res, word)
		}
	}
	return res
}

func findWords2(words []string) []string {
	line1 := "qwertyuiopQWERTYUIOP"
	line2 := "asdfghjklASDFGHJKL"
	line3 := "zxcvbnmZXCVBNM"
	var res []string
	for _, word := range words{
		b1 := strings.ContainsAny(word, line1)
		b2 := strings.ContainsAny(word, line2)
		b3 := strings.ContainsAny(word, line3)
		if (b1 && !b2 && !b3) || (!b1 && b2 && !b3) || (!b1 && !b2 && b3){
			res = append(res, word)
		}
	}
	return res
}