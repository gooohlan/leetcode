package Array

func getModifiedArray(length int, updates [][]int) []int {
	diff := make([]int, length+1)
	for _, update := range updates {
		diff[update[0]] += update[2]
		if update[1] < length {
			diff[update[1]+1] -= update[2]
		}
	}

	for i := 1; i < length; i++ {
		diff[i] += diff[i-1]
	}
	return diff[:length]
}
