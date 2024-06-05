package Everyday

func distributeCandies575(candyType []int) int {
    m := map[int]struct{}{}
    for _, c := range candyType {
        m[c] = struct{}{}
    }
    return min(len(m), len(candyType)/2)
}
