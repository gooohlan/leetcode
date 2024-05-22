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

func Test1542(t *testing.T) {
    s := "1317"
    fmt.Println(longestAwesome(s))
}
