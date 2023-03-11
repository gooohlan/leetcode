package BFS

import (
    "strconv"
    "strings"
)

// 将二维数组压缩成一维数组,将二维数组中的上下左右概念,通过索引映射可以实现,下面举例说明:
// 题目中的数组大小为2*3, 对二维数组的位置进行编号会得到一个[[0,1,2][3,4,5]]的二维数组
// 分别写出每个位置相邻元素得到
// 0 的相邻位置是 1, 3；
// 1 的相邻位置是 0, 2, 4
// 2 的相邻位置是 1, 5
// 3 的相邻位置是 0, 4
// 4 的相邻位置是 1, 3, 5
// 5 的相邻位置是 2, 4
// 在压缩后的一维数组中找到要移动位置,依次移动即可
// 比如 [[4, 1, 2], [5, 0, 3]] => 412503, 0所在的位置为4,它相邻的元素位置分别是1,3,5,一次交换得到:412053, 412530, 402513
var neighbors = [6][]int{
    {1, 3},
    {0, 2, 4},
    {1, 5},
    {0, 4},
    {1, 3, 5},
    {2, 4},
}
var target = "123450"

func slidingPuzzle(board [][]int) int {
    var start string
    for _, r := range board {
        for _, i := range r {
            start += strconv.Itoa(i)
        }
    }
    
    q := []string{start}
    seen := map[string]bool{q[0]: true}
    step := 0
    
    for len(q) != 0 {
        sz := len(q)
        for i := 0; i < sz; i++ {
            p := q[0]
            q = q[1:]
            
            if p == target {
                return step
            }
            x := strings.Index(p, "0")
            for _, y := range neighbors[x] {
                s := []byte(p)
                s[x], s[y] = s[y], s[x]
                newBoard := string(s)
                if !seen[newBoard] {
                    q = append(q, newBoard)
                    seen[newBoard] = true
                }
            }
        }
        step++
    }
    return -1
}

func slidingPuzzle2(board [][]int) int {
    var start string
    for _, r := range board {
        for _, i := range r {
            start += strconv.Itoa(i)
        }
    }
    q1 := map[string]bool{start: true}
    q2 := map[string]bool{target: true}
    seen := map[string]bool{}
    step := 0
    
    for len(q1) != 0 && len(q2) != 0 {
        if len(q1) > len(q2) {
            q1, q2 = q2, q1
        }
        
        temp := make(map[string]bool)
        for q := range q1 {
            if q2[q] {
                return step
            }
            if seen[q] {
                continue
            }
            seen[q] = true
            
            x := strings.Index(q, "0")
            for _, y := range neighbors[x] {
                s := []byte(q)
                s[x], s[y] = s[y], s[x]
                temp[string(s)] = true
            }
        }
        step++
        q2, q1 = temp, q2
    }
    return -1
}
