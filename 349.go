package main

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	var intersection []int
	mapList := make(map[int]bool)
	for _, v := range nums1 {
		mapList[v] = true
	}

	for _, v := range nums2 {
		if _, ok := mapList[v]; ok && mapList[v] {
			intersection = append(intersection, v)
			mapList[v] = false
		}
	}
	return intersection
}

func intersection2(nums1 []int, nums2 []int) []int {
	var intersection []int
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i,j := 0,0; i<len(nums1) && j <len(nums2); {
		if nums1[i] == nums2[j] {
			if intersection == nil || intersection[len(intersection)-1] != nums1[i] {
				intersection = append(intersection, nums1[i])
			}
			i++
			j++
		}else if nums1[i] > nums2[j]{
			j++
		} else {
			i++
		}
	}
	return intersection
}
