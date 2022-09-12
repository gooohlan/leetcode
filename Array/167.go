package Array

// 定义左右指针分别执行数组第一个元素和最后一个元素
// 根据题意，数组是从小到大排列的，在左指针小于右指针的前提下，用他们的和依次去与target进行比较
// 比target大时，右指针往左，和变小
// 比target小时，左指针往右，和变大
func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}
