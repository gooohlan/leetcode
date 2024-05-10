package Other

func maxCount(m int, n int, ops [][]int) int {
    arr := make([][]int, m)
    for i := range arr {
        arr[i] = make([]int, n)
    }

    for _, op := range ops {
        for i := 0; i < op[0]; i++ {
            for j := 0; j < op[1]; j++ {
                arr[i][j]++
            }
        }
    }
    max := arr[0][0]
    var count int
    for _, ints := range arr {
        for _, v := range ints {
            if v == max {
                count++
            }
        }
    }
    return count
}

func maxCount2(m int, n int, ops [][]int) int {
    for _, op := range ops {
        if m > op[0] {
            m = op[0]
        }
        if n > op[1] {
            n = op[1]
        }
    }
    return m * n
}
