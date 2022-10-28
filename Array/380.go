package Array

import "math/rand"

type RandomizedSet struct {
	nums     []int
	varIndex map[int]int
}

func Constructor380() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.varIndex[val]; ok {
		return false
	}
	r.varIndex[val] = len(r.nums)
	r.nums = append(r.nums, val)
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	index, ok := r.varIndex[val]
	if !ok {
		return false
	}

	n := len(r.nums) - 1
	r.nums[index] = r.nums[n]         // 将末尾数放到移除数的位置
	r.varIndex[r.nums[index]] = index // 更新末尾数索引
	r.nums = r.nums[:n]               // 移除最后一个数
	delete(r.varIndex, val)           // 删除索引
	return true
}

func (r *RandomizedSet) GetRandom() int {
	return r.nums[rand.Intn(len(r.nums))]
}
