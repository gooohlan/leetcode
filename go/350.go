package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) { // 遍历较小的数组生成map
		intersect(nums2, nums1)
	}
	var intersection []int
	mapList := make(map[int]int)
	for _, v := range nums1 {
		mapList[v] ++
	}

	for _, v := range nums2 {
		if _, ok := mapList[v]; ok && mapList[v] > 0 {
			intersection = append(intersection, v)
			mapList[v] --
		}
	}
	return intersection
}

func intersect2(nums1 []int, nums2 []int) []int {
	var intersection []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			intersection = append(intersection, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return intersection
}
