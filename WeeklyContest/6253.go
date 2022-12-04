package WeeklyContest

import "strings"

func isCircularSentence(sentence string) bool {
	strs := strings.Split(sentence, " ")
	prev := sentence[len(sentence)-1:]
	for _, str := range strs {
		if string(str[0]) != prev {
			return false
		}
		prev = str[len(str)-1:]
	}
	return true
}
