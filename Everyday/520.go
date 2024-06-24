package Everyday

import "unicode"

func detectCapitalUse(word string) bool {
    count := 0
    for _, s := range word {
        if unicode.IsUpper(s) {
            count++
        }
    }
    
    return count == 0 || count == len(word) || count == 1 && unicode.IsUpper(rune(word[0]))
}
