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
    nums := []int{1, 2, 3, 6}
    fmt.Println(canPartition(nums))
}

func Test518(t *testing.T) {
    nums := []int{1, 2, 5}
    fmt.Println(change(10, nums))
    fmt.Println(change2(10, nums))
}

func Test494(t *testing.T) {
    nums := []int{1, 1, 1, 1, 1}
    fmt.Println(findTargetSumWays(nums, 3))
    fmt.Println(findTargetSumWaysDFS(nums, 3))
    fmt.Println(findTargetSumWaysDP(nums, 3))
    fmt.Println(findTargetSumWaysDP2(nums, 3))
}

func Test64(t *testing.T) {
    grid := [][]int{
        []int{1, 2, 3},
        []int{4, 5, 6},
    }
    fmt.Println(minPathSum(grid))
    fmt.Println(minPathSumDP(grid))
    fmt.Println(minPathSumDP2(grid))
}

func Test174(t *testing.T) {
    dungeon := [][]int{
        []int{-2, -3, 3},
        []int{-5, -10, 1},
        []int{10, 30, -5},
    }
    dungeon2 := [][]int{
        []int{-1, -5},
        []int{-4, -3},
    }
    
    fmt.Println(calculateMinimumHP(dungeon2))
    fmt.Println(calculateMinimumHPDP(dungeon))
    fmt.Println(calculateMinimumHPDP2(dungeon))
}

func Test514(t *testing.T) {
    fmt.Println(findRotateSteps("godding", "godding"))
    fmt.Println(findRotateStepsDP1("godding", "godding"))
    fmt.Println(findRotateStepsDP2("godding", "godding"))
    fmt.Println(findRotateStepsDP2_2("godding", "gd"))
}

func Test787(t *testing.T) {
    edges := [][]int{
        []int{0, 1, 100},
        []int{1, 2, 100},
        []int{0, 2, 500},
    }
    
    fmt.Println(findCheapestPrice(3, edges, 0, 2, 0))
}
