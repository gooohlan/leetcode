package DataStructure

func trailingZeroes(n int) (res int) {
    for n != 0 {
        res += n / 5
        n /= 5
    }
    return res
}
