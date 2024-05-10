package Everyday

func minimumRefill(plants []int, capacityA int, capacityB int) int {
    res := 0
    n := len(plants)
    i, j := 0, n-1
    va, vb := capacityA, capacityB
    for i < j {
        if va < plants[i] {
            va = capacityA
            res++
        }
        va -= plants[i]

        i++

        if vb < plants[j] {
            res++
            vb = capacityB
        }
        vb -= plants[j]
        j--
    }

    if i == j && max(va, vb) < plants[i] {
        res++
    }
    return res
}
