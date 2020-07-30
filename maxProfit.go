package main

import "fmt"

func main()  {
	fmt.Println(maxProfit([]int{2,1,2,0,1}))
}


func maxProfit(prices []int) int {
	var i, max int

	start := -1

	l := len(prices)

	//var temp []int
	for i = 0; i < l; i++ {
		if i < l-1 && prices[i] < prices[i+1] {
			if start == -1 {
				start = prices[i]
			}

			if start > prices[i] {
				start = prices[i]
			}
		} else {

			if start >= 0 {
				max += prices[i] - start
			}
			start = -1
		}
	}

	return max
}

