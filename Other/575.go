package Other

import "sort"

func distributeCandies(candyType []int) int {
    sort.Ints(candyType)
    count := 1
    for i := 1; i < len(candyType) && count < len(candyType)/2; i++ {
        if candyType[i] != candyType[i-1] {
            count++
        }
    }
    return count
}

func distributeCandiesMap(candyType []int) int {
    mapList := make(map[int]struct{}, 0)
    for _, v := range candyType {
        mapList[v] = struct{}{}
    }
    if len(mapList) < len(candyType)/2 {
        return len(mapList)
    }
    return len(candyType) / 2
}
