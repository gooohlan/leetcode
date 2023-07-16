package WeeklyContest

func sumOfSquares(nums []int) int {
    n := len(nums)
    sum := 0
    for i := 1; i < n+1; i++ {
        if n%i == 0 {
            sum = sum + nums[i-1]*nums[i-1]
        }
    }
    return sum
}
