package main

import "fmt"

func main() {
	fmt.Println(strStr("hello world", "ll"))
}

/**
  字符串查找

  o(n^2)
 */
func strStr(haystack string, needle string) int  {
	if len(needle) == 0 {
		return 0
	}

	var i, j int

	// 循环外层被搜索字符串
	// len(haystack) -len(needle) + 1 => "world", "rl" => 祝需要搜索 worl 即可
	for i = 0; i< len(haystack) -len(needle) + 1; i++ {
		// 循环查找字符串
		for j = 0; j < len(needle); j++ {
			// 一旦又不相等的字符 退出循环
			if haystack[i+j] != needle[j] {
				break
			}
		}
		// break 之后 判断 len(needle) == j
		// 相等说明 needle 循环结束
		// 说明 needle 已找到
		// 可以返回 i 位置
		if len(needle) == j {
			return i
		}
	}

	return -1
}