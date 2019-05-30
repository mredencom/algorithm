package main

import (
	"fmt"
	"strings"
)

//leetcode-cn 网站 28 实现strStr
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	s := strStr("hello", "ll")
	fmt.Println(s)
}
