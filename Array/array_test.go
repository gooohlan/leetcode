package Array

import (
    "fmt"
    "testing"
)

func Test303(t *testing.T) {
    nums := Constructor([]int{-2, 0, 3, -5, 2, -1})
    
    fmt.Println(nums.SumRange(0, 2))
    fmt.Println(nums.SumRange(2, 5))
    fmt.Println(nums.SumRange(0, 5))
}

func Test304(t *testing.T) {
    nums := [][]int{
        []int{3, 0, 1, 4, 2},
        []int{5, 6, 3, 2, 1},
        []int{1, 2, 0, 1, 5},
        []int{4, 1, 0, 1, 7},
        []int{1, 0, 3, 0, 5}}
    matrix := ConstructorMatrix(nums)
    fmt.Println(matrix.SumRegion(2, 1, 4, 3))
    fmt.Println(matrix.SumRegion(1, 1, 2, 2))
    fmt.Println(matrix.SumRegion(1, 2, 2, 4))
}

func Test370(t *testing.T) {
    updates := [][]int{
        []int{1, 3, 2},
        []int{2, 4, 3},
        []int{0, 2, -2},
    }
    
    fmt.Println(getModifiedArray(5, updates))
}

func Test1109(t *testing.T) {
    bookings := [][]int{
        []int{1, 2, 10},
        []int{2, 2, 15},
    }
    fmt.Println(corpFlightBookings(bookings, 2))
}

func Test1094(t *testing.T) {
    trips := [][]int{
        []int{9, 0, 1},
        []int{3, 3, 7},
    }
    fmt.Println(carPooling(trips, 4))
}

func Test151(t *testing.T) {
    str := "     xxxx  cccc   aaaa   cd sc "
    fmt.Println(reverseWords(str))
}

func Test48(t *testing.T) {
    matrix := [][]int{
        []int{1, 1, 1},
        []int{2, 2, 2},
        []int{3, 3, 3},
    }
    rotate(matrix)
    fmt.Println(matrix)
}

func Test54(t *testing.T) {
    matrix := [][]int{
        []int{1, 2, 3, 4},
        []int{5, 6, 7, 8},
        []int{9, 10, 11, 12},
    }
    
    fmt.Println(spiralOrder(matrix))
}

func Test59(t *testing.T) {
    matrix := generateMatrix(2)
    fmt.Println(matrix)
}

func Test528(t *testing.T) {
    s := Constructor528([]int{1, 3, 2})
    numMap := make(map[int]int)
    for i := 0; i < 10000; i++ {
        numMap[s.PickIndex()+1]++
    }
    fmt.Println(numMap)
}

func Test870(t *testing.T) {
    nums1 := []int{12, 24, 8, 32}
    nums2 := []int{13, 25, 32, 11}
    fmt.Println(advantageCount(nums1, nums2))
}

func Test380(t *testing.T) {
    constructor := Constructor380()
    fmt.Println(constructor.Insert(1))
    fmt.Println(constructor.Insert(2))
    fmt.Println(constructor.Insert(3))
    fmt.Println(constructor.Insert(4))
    fmt.Println(constructor.Insert(5))
    fmt.Println(constructor.Insert(5))
    fmt.Println(constructor.Remove(5))
    fmt.Println(constructor.Insert(5))
}

func Test710(t *testing.T) {
    constructor := Constructor710(5, []int{1, 4})
    fmt.Println(constructor.Pick())
}

func Test1081(t *testing.T) {
    str := "cbacdcbc"
    fmt.Println(smallestSubsequence(str))
}

func Test316(t *testing.T) {
    str := "cbacdcbc"
    fmt.Println(removeDuplicateLetters(str))
}

func Test398(t *testing.T) {
    solution := Constructor398([]int{1, 2, 3, 3, 3})
    fmt.Println(solution.Pick(3))
    fmt.Println(solution.Pick(1))
    fmt.Println(solution.Pick(3))
}

func Test659(t *testing.T) {
    nums := []int{1, 2, 3, 3, 4, 5}
    fmt.Println(isPossible(nums))
    getPossible(nums)
}

func Test969(t *testing.T) {
    arr := []int{3, 2, 4, 1}
    
    fmt.Println(pancakeSort(arr))
}
