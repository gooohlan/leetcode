package WeeklyContest

func differenceOfSum(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
        for num > 0 {
            sum -= num % 10
            num /= 10
        }
    }
    return sum
}
