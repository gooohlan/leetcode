package Other

func findPoisonedDuration(timeSeries []int, duration int) int {
    if len(timeSeries) == 0 {
        return 0
    }
    var total int
    for i := 1; i < len(timeSeries); i++ {
        if timeSeries[i]-timeSeries[i-1] < duration {
            total += timeSeries[i] - timeSeries[i-1]
        } else {
            total += duration
        }
    }
    return total + duration
}
