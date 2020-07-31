package main

import "fmt"

func main()  {
	// fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))

	removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4})
}

func removeElement(nums []int, val int) int {

	for i:=0;i<len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return len(nums)
}

func removeDuplicates(nums []int) int {
	for i:=0;i+1<len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	fmt.Println(nums)
	return len(nums)
}