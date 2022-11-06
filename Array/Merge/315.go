package Merge

import "fmt"

// 在归并排序的基础上,我们在和的时候做文章
// 在合并时,左侧元素小于右侧元素时,此时右侧元素到中间为止的距离,是比比他小的元素个数
// 举个例子, L = [1, 1, 3, 5]; R = [2, 4, 6, 7]
// 假设此时左边指针指向5, 右边指向6, 此时就可以得出右边比5小的个数为 2, 4
// 每当执行 nums[p] = temp[i] 时，就可以确定 temp[i] 这个元素后面比它小的元素个数为 j - mid - 1
// 这是对于一个子数组,这个子数组之后还会被合并,再合并时我们将个数叠加即可
type Pair struct {
	Val, Index int // val记录元素值, index记录索引
}

var tmp []Pair
var countList []int

func countSmaller(nums []int) []int {
	tmp = make([]Pair, len(nums))
	countList = make([]int, len(nums))
	arr := make([]Pair, len(nums))
	for i, num := range nums {
		arr[i] = Pair{
			Val:   num,
			Index: i,
		}
	}
	sort(arr, 0, len(nums)-1)
	return countList
}

func sort(arr []Pair, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	sort(arr, l, mid)
	sort(arr, mid+1, r)
	merge315(arr, l, mid, r)
}

func merge315(arr []Pair, l, mid, r int) {
	i, j := l, mid+1
	cnt := l

	for i := l; i <= r; i++ {
		tmp[i] = arr[i]
	}

	for i <= mid && j <= r {
		if tmp[i].Val <= tmp[j].Val {
			arr[cnt] = tmp[i]
			i++
			// 记录右侧比tmp[i]的个数, arr[cnt].Index保留这原始下标
			countList[arr[cnt].Index] += j - mid - 1
			fmt.Println(countList)
		} else {
			arr[cnt] = tmp[j]
			j++
		}
		cnt++
	}

	for i <= mid {
		arr[cnt] = tmp[i]
		i++
		// 记录右侧比tmp[i]的个数, arr[cnt].Index保留这原始下标
		countList[arr[cnt].Index] += j - mid - 1
		fmt.Println(countList)
		cnt++
	}

	for j <= r {
		arr[cnt] = tmp[j]
		j++
		cnt++
	}
}
