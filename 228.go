package main

import (
	"log"
	"strconv"
)

func summaryRanges(nums []int) []string {
	var rStr []string
	if len(nums) == 0 {
		return rStr
	}
	var head, tail int
	head, tail = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != tail+1 {
			if head == tail {
				rStr = append(rStr, strconv.Itoa(head))
			} else {
				rStr = append(rStr, strconv.Itoa(head)+"->"+strconv.Itoa(tail))
			}
			head = nums[i]
			tail = nums[i]
		}else{

			tail = nums[i]
		}
	}
	if head == tail {
		rStr = append(rStr, strconv.Itoa(head))
	} else {
		rStr = append(rStr, strconv.Itoa(head)+"->"+strconv.Itoa(tail))
	}
	return rStr
}

func main() {
	str := summaryRanges1([]int{0, 2, 3, 4, 6, 8, 9})
	log.Println(str)
}

func summaryRanges1(nums []int) []string {
	var rStr []string
	for i,j :=0,0; j<len(nums);{
		if j+1 < len(nums) && nums[j+1] == nums[j]+1{
			j++
			continue
		}
		if i==j {
			rStr = append(rStr, strconv.Itoa(nums[i]))
		}else{
			rStr = append(rStr, strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j]))
		}
		i = j+1
		j++
	}
	return rStr
}
