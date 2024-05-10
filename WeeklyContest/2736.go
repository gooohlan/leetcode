package WeeklyContest

import "sort"

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
    type pair struct {
        x, y int
    }
    
    // x从小到大排序
    pairList := make([]pair, len(nums1))
    for i, x := range nums1 {
        pairList[i] = pair{x, nums2[i]}
    }
    sort.Slice(pairList, func(i, j int) bool { return pairList[i].x < pairList[j].x })
    
    for i := range queries {
        queries[i] = append(queries[i], i) // 将下标加入queries
    }
    // 查询从大到小排序
    sort.Slice(queries, func(i, j int) bool { return queries[i][0] > queries[j][0] })
    
    ans := make([]int, len(queries))
    var stack []pair
    i := len(pairList) - 1
    for _, query := range queries {
        for i >= 0 && pairList[i].x >= query[0] {
            for len(stack) > 0 && stack[len(stack)-1].y <= pairList[i].x+pairList[i].y { // 栈顶小于当前值,弹出
                stack = stack[:len(stack)-1]
            }
            if len(stack) == 0 || stack[len(stack)-1].x < pairList[i].y { // 栈顶大于当前值,压入
                stack = append(stack, pair{pairList[i].y, pairList[i].x + pairList[i].y})
            }
            i--
        }
        
        j := sort.Search(len(stack), func(i int) bool { return stack[i].x >= query[1] })
        if j < len(stack) {
            ans[query[2]] = stack[j].y
        } else {
            ans[query[2]] = -1
        }
    }
    return ans
}
