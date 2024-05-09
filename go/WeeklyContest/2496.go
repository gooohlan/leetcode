package WeeklyContest

import (
	"math"
	"strconv"
)

func maximumValue(strs []string) int {
	max := math.MinInt
	for _, str := range strs {
		value := isNums(str)
		if max < value {
			max = value
		}
	}

	return max
}

func isNums(str string) int {
	sum := 0
	for _, s := range str {
		sum += int(s)
	}
	if sum > 59*len(str) {
		return len(str)
	}
	i, _ := strconv.Atoi(str)
	return i
}
