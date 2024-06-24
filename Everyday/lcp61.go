package Everyday

func temperatureTrend(temperatureA []int, temperatureB []int) int {
    sum := 0
    ans := 0
    for i := 1; i < len(temperatureA); i++ {
        if isCompare(temperatureA[i], temperatureA[i-1]) == isCompare(temperatureB[i], temperatureB[i-1]) {
            sum++
            ans = max(ans, sum)
        } else {
            sum = 0
        }
    }
    return ans
}

func isCompare(a, b int) int {
    if a > b {
        return 1
    } else if a == b {
        return 0
    }
    return -1
}
