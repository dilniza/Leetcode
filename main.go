package main

import (
	"fmt"
	"leet/leetcode"
)

func main() {
	var nums []int
	nums = []int{1,11,10}
	k := 3
	res := leetcode.ResultsArray(nums, k)
	fmt.Println(res)

	nums = []int{1,2,9,8}
	k = 2
	res = leetcode.ResultsArray(nums, k)
	fmt.Println(res)
}