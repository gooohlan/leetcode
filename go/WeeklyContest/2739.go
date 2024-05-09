package WeeklyContest

func distanceTraveled1(mainTank int, additionalTank int) int {
    var res int
    for mainTank >= 5 {
        mainTank -= 5
        res += 5 * 10
        if additionalTank > 0 {
            additionalTank--
            mainTank++
        }
    }
    res += mainTank * 10
    return res
}

func distanceTraveled2(mainTank int, additionalTank int) int {
    var res int
    for mainTank >= 5 {
        t := mainTank / 5
        mainTank %= 5
        res += t * 50
        t = min(t, additionalTank)
        additionalTank -= t
        mainTank += t
    }
    res += mainTank * 10
    return res
}

// 每次消耗5升,获取1升,就相当于消耗4升,这样相当于可以消耗mainTank/4升
// 但是mainTank是4的倍数,最后一次消耗是的不到补充的,所以实际消耗次数为(mainTank-1)/4
// 补充的油不能超过备用油箱中的油
func distanceTraveled3(mainTank int, additionalTank int) int {
    return (min((mainTank-1)/4, additionalTank) + mainTank) * 10
}
