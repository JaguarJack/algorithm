package main

import (
	"fmt"
)

func main()  {
	fmt.Println(plusOne([]int{9, 9}))
}

func plusOne(digits []int) []int {
	l := len(digits) - 1

	for l >= 0 {
		if (digits[l] + 1) == 10 {
			digits[l] = 0
			if l == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[l] = digits[l] + 1
			break
		}
		l--
	}

	return digits
}