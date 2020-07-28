package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(longestCommonPrefix([]string{ "aa", "aa"}))
}

func longestCommonPrefix(strs []string)  string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	prefix :=  strs[0]

	for _, str := range strs {
		prefix = findPrefix(prefix, str)
		if prefix == "" {
			return prefix
		}
	}

	return prefix
}

func findPrefix(first, second string) string {
	l := len(second)

	for l > 0 {
		second = second[:l]

		if strings.Index(first, second) == 0 {
			return second
		}
		l--
	}

	return ""
}