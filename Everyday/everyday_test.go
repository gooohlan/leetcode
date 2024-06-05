package Everyday

import (
    "fmt"
    "testing"
)

func Test2391(t *testing.T) {
    garbage := []string{"G", "P", "GP", "GG"}
    travel := []int{2, 4, 3}
    fmt.Println(garbageCollection2(garbage, travel))
}

func Test2589(t *testing.T) {
    tasks := [][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}}
    fmt.Println(findMinimumTime(tasks))
}

func Test1542(t *testing.T) {
    s := "1317"
    fmt.Println(longestAwesome(s))
}

func Test826(t *testing.T) {
    difficulty := []int{2, 4, 6, 8, 10}
    profit := []int{10, 20, 30, 40, 50}
    worker := []int{4, 5, 6, 7}
    maxProfitAssignment(difficulty, profit, worker)
}

func Test2831(t *testing.T) {
    nums := []int{1, 3, 2, 3, 1, 3, 4}
    k := 3
    fmt.Println(longestEqualSubarray(nums, k))
}

func Test1738(t *testing.T) {
    matrix := [][]int{
        []int{5, 2},
        []int{1, 6},
        []int{8, 3},
    }
    fmt.Println(kthLargestValue(matrix, 3))
}

func Test2951(t *testing.T) {
    arr := []int{1, 4, 3, 8, 5}
    fmt.Println(findPeaks(arr))
}

func Test2982(t *testing.T) {
    s := "aada"
    fmt.Println(maximumLength(s))
}

func Test2928(t *testing.T) {
    n := 5
    limit := 2
    fmt.Println(distributeCandies(n, limit))
}
