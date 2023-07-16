package WeeklyContest

func minimumIndex(nums []int) int {
    freq := map[int]int{}
    x := nums[0]
    for _, num := range nums {
        freq[num]++
        if freq[num] > freq[x] {
            x = num
        }
    }
    
    total := freq[x]
    freq1 := 0
    for i, num := range nums {
        if num == x {
            freq1++
        }
        if freq1*2 > i+1 && (total-freq1)*2 > len(nums)-i-1 {
            return i
        }
    }
    return -1
}
