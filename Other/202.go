package Other

func isHappy(n int) bool {
    m := make(map[int]bool)
    for ; n != 1 && !m[n]; n, m[n] = step(n), true {
    }
    return n == 1
}

func step(n int) int {
    sum := 0
    for n > 0 {
        sum += (n % 10) * (n % 10)
        n /= 10
    }
    return sum
}

func isHappyList(n int) bool {
    slow, fast := n, step(n)
    for ; fast != 1 && slow != fast; {
        slow = step(slow)
        fast = step(step(fast))
    }
    return fast == 1
}
