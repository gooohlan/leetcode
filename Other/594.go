package Other

func findLHS(nums []int) int {
    mapList := make(map[int]int)
    var res int
    for _, num := range nums {
        mapList[num]++
        if _, ok := mapList[num-1]; ok && res < (mapList[num-1]+mapList[num]) {
            res = mapList[num-1] + mapList[num]
        }
        if _, ok := mapList[num+1]; ok && res < (mapList[num+1]+mapList[num]) {
            res = mapList[num+1] + mapList[num]
        }
    }
    return res
}
