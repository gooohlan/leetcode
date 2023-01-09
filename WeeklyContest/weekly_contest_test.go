package WeeklyContest

import (
    "fmt"
    "testing"
)

func Test6245(t *testing.T) {
    fmt.Println(pivotInteger(1))
}

func Test6246(t *testing.T) {
    fmt.Println(appendCharacters("abcde", "a"))
}

func Test6247(t *testing.T) {
    head := &ListNode{
        Val: 5,
        Next: &ListNode{
            Val: 2,
            Next: &ListNode{
                Val: 13,
                Next: &ListNode{
                    Val: 3,
                    Next: &ListNode{
                        Val:  8,
                        Next: nil,
                    },
                },
            },
        },
    }
    
    // fmt.Println(removeNodes(head))
    head = removeNodes2(head)
    fmt.Println(1)
}

func Test6248(t *testing.T) {
    nums := []int{5, 1, 7, 4, 2, 6, 3}
    fmt.Println(countSubarrays(nums, 4))
}

func Test6255(t *testing.T) {
    roads := [][]int{
        []int{1, 2, 9},
        []int{2, 3, 6},
        []int{2, 4, 5},
        []int{1, 4, 7},
    }
    
    fmt.Println(minScore(4, roads))
}

func Test6256(t *testing.T) {
    edges := [][]int{
        []int{1, 2},
        []int{1, 4},
        []int{1, 5},
        []int{2, 6},
        []int{2, 3},
        []int{4, 6},
    }
    fmt.Println(magnificentSets(6, edges))
}

func Test6261(t *testing.T) {
    strs := []string{"1", "01", "001", "0001"}
    fmt.Println(maximumValue(strs))
}

func Test6262(t *testing.T) {
    vals := []int{
        1, 2, 3, -4, 10, -10, 20,
    }
    edges := [][]int{
        []int{0, 1},
        []int{1, 2},
        []int{1, 3},
        []int{3, 4},
        []int{3, 5},
        []int{3, 6},
    }
    fmt.Println(maxStarSum(vals, edges, 2))
}

func Test6263(t *testing.T) {
    fmt.Println(maxJump([]int{0, 2, 5, 6, 15}))
}

func Test6257(t *testing.T) {
    grid := [][]int{
        []int{10},
    }
    
    fmt.Println(deleteGreatestValue(grid))
}

func Test6258(t *testing.T) {
    nums := []int{4, 3, 6, 16, 8, 2}
    fmt.Println(longestSquareStreak(nums))
}

func Test6259(t *testing.T) {
    loc := Constructor(10)
    fmt.Println(loc.Allocate(1, 1))
    fmt.Println(loc.Allocate(1, 2))
    fmt.Println(loc.Allocate(1, 3))
    fmt.Println(loc.Free(2))
    fmt.Println(loc.Allocate(3, 4))
    fmt.Println(loc.Allocate(1, 1))
    fmt.Println(loc.Allocate(1, 1))
    fmt.Println(loc.Free(1))
    fmt.Println(loc.Allocate(10, 2))
    fmt.Println(loc.Free(7))
}

func Test6260(t *testing.T) {
    grid := [][]int{
        []int{1, 2, 3},
        []int{2, 5, 7},
        []int{3, 5, 1},
    }
    queries := []int{5, 6, 2}
    fmt.Println(maxPoints2(grid, queries))
}

func Test6264(t *testing.T) {
    fmt.Println(minimumTotalCost([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}))
}

func Test6265(t *testing.T) {
    words := []string{"aba", "aabb", "ab", "bbaa", "aabc"}
    fmt.Println(similarPairs(words))
    fmt.Println(similarPairs2(words))
}

func Test6266(t *testing.T) {
    fmt.Println(smallestValue(15))
}

func Test6267(t *testing.T) {
    edges := [][]int{
        []int{1, 2},
        []int{1, 4},
        []int{1, 3},
    }
    fmt.Println(isPossible(4, edges))
}

func Test6269(t *testing.T) {
    words := []string{"a", "b", "c", "c", "d", "e", "f", "g", "h", "j"}
    fmt.Println(closetTarget(words, "c", 8))
}

func Test6270(t *testing.T) {
    fmt.Println(takeCharacters("aabaaaacaabc", 2))
}

func Test6271(t *testing.T) {
    fmt.Println(maximumTastiness([]int{13, 5, 1, 8, 21, 2}, 3))
    fmt.Println(maximumTastiness2([]int{13, 5, 1, 8, 21, 2}, 3))
}

func Test6272(t *testing.T) {
    nums := []int{1, 2, 3, 4}
    fmt.Println(countPartitions(nums, 4))
}

func Test6273(t *testing.T) {
    num := 1248
    fmt.Println(countDigits(num))
}

func Test6274(t *testing.T) {
    nums := []int{2, 4, 3, 7, 10, 6}
    fmt.Println(distinctPrimeFactors(nums))
}

func Test6275(t *testing.T) {
    fmt.Println(minimumPartition("165462", 60))
}

func Test7276(t *testing.T) {
    fmt.Println(closestPrimes(10, 19))
}

func Test6283(t *testing.T) {
    fmt.Println(maximumCount2([]int{-3, -2, -1, 0, 0, 1, 2}))
}

func Test6285(t *testing.T) {
    fmt.Println(maxKelements([]int{10, 10, 10, 10, 10}, 5))
}

func Test6284(t *testing.T) {
    fmt.Println(isItPossible("abcc", "aab"))
}
