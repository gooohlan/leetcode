package Other

import (
    "sort"
    "strconv"
)

// 将原位置存入Hash
// 排序
// 循环原数组获得排名
// 1:0 3:1 2:2 3:4
// 1 2 3 4
func findRelativeRanks(score []int) []string {
    mapList := make(map[int]int, len(score))
    for k, v := range score {
        mapList[v] = k
    }
    sort.Ints(score)
    res := make([]string, len(score))
    for i := 0; i < len(score); i++ {
        if len(score)-i == 1 {
            res[mapList[score[i]]] = "Gold Medal"
            continue
        }
        if len(score)-i == 2 {
            res[mapList[score[i]]] = "Silver Medal"
            continue
        }
        if len(score)-i == 3 {
            res[mapList[score[i]]] = "Bronze Medal"
            continue
        }
        res[mapList[score[i]]] = strconv.Itoa(len(score) - i)
    }

    return res
}
