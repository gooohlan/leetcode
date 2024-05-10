package DataStructure

func bulbSwitch(n int) int {
    var ans int
    for i := 1; i*i <= n; i++ {
        ans++
    }
    return ans
}
