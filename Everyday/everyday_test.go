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
