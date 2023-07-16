package WeeklyContest

func sumOfSquares(nums []int) int {
    var res int
    for i, num := range nums {
        if len(nums)%(i+1) == 0 {
            res += num * num
        }
    }
    return res
}
