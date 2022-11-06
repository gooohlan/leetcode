package Merge

// 先构件前缀和数组,归并是进行处理
// 给定两个升序排列的数组n1, n2分别找出下标1,j满足lower <= n2[j]-n1[i] <= upper
// 在n2种维护两个指针l,r都指向n2初始位置
// l右移,直到n1[0]+n2[l]>=lower;r右移,直到n1[0]+n2[r]<upper
// 此时在[l,r)这个区间种所有的下标j,都满足lower <= n2[j]-n1[0] <= upper
// 因为n1也是有序数组, 所以n1后移时,l和r跟着后移就可以了
// func countRangeSum(nums []int, lower int, upper int) int {
// 	tmp := make([]int, len(nums))
//
// }
