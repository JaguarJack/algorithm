package main

import (
	"fmt"
	"sort"
)

/**
给定两个数组，编写一个函数来计算它们的交集。

输入: nums1 = [1,2,2,1], nums2 = [2,2]

输出: [2,2]

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]

输出: [4,9]

 */
func main()  {

	fmt.Println(_intersect([]int{1,2,2,1,4,5,7}, []int{2,2,5,3}))
}

func _intersect(num1, num2 []int) []int {
	sort.Ints(num1)
	sort.Ints(num2)

	var i,j int

	var temp []int

	for i < len(num1) && j <len(num2) {
		if num1[i] == num2[j] {
			temp = append(temp, num1[i])
			i++
			j++
		} else if num1[i] < num2[j] {
			i++
		} else {
			j++
		}
	}

	return temp
}

func intersect(num1, num2 []int) []int  {
	m := map[int]int{}

	for _,v := range num1 {
		m[v] += 1
	}

	k := 0

	for _, v := range num2 {
		if m[v] > 0 {
			m[v] -= 1
			num2[k] = v
			k++
		}
	}

	return num2[0:k]
}