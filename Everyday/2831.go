package Everyday

func longestEqualSubarray(nums []int, k int) int {
    pos := make(map[int][]int)
    
    for i, num := range nums {
        pos[num] = append(pos[num], i)
    }
    
    res := 0
    for _, p := range pos {
        left := 0
        for right := range p {
            for p[right]-p[left]-right+left > k {
                left++
            }
            res = max(res, right-left+1)
        }
    }
    return res
}

func longestEqualSubarray2(nums []int, k int) int {
    pos := make(map[int][]int)
    
    for i, num := range nums {
        pos[num] = append(pos[num], i-len(pos[num]))
    }
    
    res := 0
    for _, p := range pos {
        left := 0
        for right := range p {
            for p[right]-p[left] > k {
                left++
            }
            res = max(res, right-left+1)
        }
    }
    return res
}
