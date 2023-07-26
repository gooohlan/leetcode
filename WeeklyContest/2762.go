package WeeklyContest

func continuousSubarrays(nums []int) int64 {
    cnt := map[int]int{}
    left := 0
    var res int64
    for right, num := range nums {
        cnt[num]++
        for {
            maxN, minN := num, num
            for k := range cnt {
                maxN = max(maxN, k)
                minN = min(minN, k)
            }
            if maxN-minN <= 2 {
                break
            }
            
            y := nums[left]
            cnt[y]--
            if cnt[y] == 0 {
                delete(cnt, y)
            }
            left++
        }
        res += int64(right - left + 1)
    }
    return res
}
