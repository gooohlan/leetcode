package WeeklyContest

func distinctPrimeFactors(nums []int) int {
    set := map[int]struct{}{}
    
    for _, num := range nums {
        for i := 2; i*i <= num; i++ {
            if num%i == 0 {
                set[i] = struct{}{}
                num /= i
                for num%i == 0 {
                    num /= i
                }
            }
        }
        if num > 1 {
            set[num] = struct{}{}
        }
    }
    return len(set)
}
