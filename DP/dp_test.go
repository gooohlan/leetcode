package DP

import (
    "fmt"
    "testing"
)

func TestFib(t *testing.T) {
    fmt.Println(fib1(20))
    fmt.Println(fib2(20))
    fmt.Println(fib3(20))
    fmt.Println(fib4(20))
}

func TestCoinChange(t *testing.T) {
    cois := []int{1, 2, 5}
    amount := 11
    fmt.Println(coinChange(cois, amount))
    fmt.Println(coinChange1(cois, amount))
    fmt.Println(coinChange2(cois, amount))
}

func Test198(t *testing.T) {
    nums := []int{1, 5, 2, 4, 1, 1, 6}
    fmt.Println(rob(nums))
    fmt.Println(rob1(nums))
    fmt.Println(rob2(nums))
    fmt.Println(rob3(nums))
}

func Test213(t *testing.T) {
    nums := []int{1, 2, 3, 1}
    fmt.Println(rob213(nums))
}

func Test337(t *testing.T) {
    node := &TreeNode{
        Val:   1,
        Left:  &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, &TreeNode{Val: 5}}},
        Right: &TreeNode{9, nil, nil},
    }
    fmt.Println(rob337(node))
}
func Test410(t *testing.T) {
    nums := []int{7, 2, 5, 10, 8}
    fmt.Println(splitArrayDP(nums, 2))
}

func Test300(t *testing.T) {
    nums := []int{0, 1, 0, 3, 2, 3}
    fmt.Println(lengthOfLIS2(nums))
}

func Test354(t *testing.T) {
    envelopes := [][]int{
        []int{4, 5},
        []int{4, 6},
        []int{6, 7},
        []int{2, 3},
        []int{1, 1},
    }
    fmt.Println(maxEnvelopes(envelopes))
}

func Test72(t *testing.T) {
    fmt.Println(minDistance("horse", "ros"))
    fmt.Println(minDistanceDP("horse", "ros"))
}

func Test931(t *testing.T) {
    matrix := [][]int{
        {2, 1, 3},
        {6, 5, 4},
        {7, 8, 9},
    }
    fmt.Println(minFallingPathSum(matrix))
    fmt.Println(minFallingPathSumDP(matrix))
}

func Test53(t *testing.T) {
    nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
    fmt.Println(maxSubArrayDP(nums))
    fmt.Println(maxSubArray(nums))
    fmt.Println(maxSubArrayPreSum(nums))
}

func Test1143(t *testing.T) {
    fmt.Println(longestCommonSubsequence("abcde", "ace"))
    fmt.Println(longestCommonSubsequenceDP("abcde", "ace"))
}

func Test583(t *testing.T) {
    fmt.Println(minDistance583DP("sea", "eat"))
    fmt.Println(minDistance583DP("leetcode", "etco"))
}

func Test712(t *testing.T) {
    fmt.Println(minimumDeleteSumDP("sea", "eat"))
    fmt.Println(minimumDeleteSumDP("delete", "leet"))
}

func Test516(t *testing.T) {
    fmt.Println(longestPalindromeSubseq("bbbab"))
    fmt.Println(longestPalindromeSubseq("cbbd"))
    fmt.Println(longestPalindromeSubseqDP("bbbab"))
    fmt.Println(longestPalindromeSubseqDP("cbbd"))
}

func Test1312(t *testing.T) {
    fmt.Println(minInsertions("mbadm"))
    fmt.Println(minInsertions("zzazz"))
    fmt.Println(minInsertions("leetcode"))
}

func Test416(t *testing.T) {
    nums := []int{1, 2, 3, 5}
    fmt.Println(canPartition(nums))
}
