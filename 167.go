package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}
		if sum < target {
			l++
		} else {
			r--
		}
	}
	return []int{-1, -1}
}
