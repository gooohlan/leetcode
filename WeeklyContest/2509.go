package WeeklyContest

// 寻找两个节点的公共父节点,向上遍历次数+1既是结果
func cycleLengthQueries(n int, queries [][]int) []int {
    res := make([]int, len(queries))
    for i, q := range queries {
        res[i] = 1
        for a, b := q[0], q[1]; a != b; {
            if a > b {
                a /= 2
            } else {
                b /= 2
            }
            res[i]++
        }
    }
    return res
}
