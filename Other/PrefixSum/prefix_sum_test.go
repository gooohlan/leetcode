package PrefixSum

import "testing"

func Test1177(t *testing.T) {
    canMakePaliQueries("abcda", [][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1}})
}
