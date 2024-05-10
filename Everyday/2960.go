package Everyday

func countTestedDevices(batteryPercentages []int) int {
    res := 0

    for _, percentage := range batteryPercentages {
        if percentage > res {
            res++
        }
    }
    return res
}
