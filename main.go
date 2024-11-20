package main

import (
	"fmt"
	"leet/leetcode"
)

func main() {
	s := "abc"
	k := 2
	res := leetcode.TakeCharacters(s, k)
	fmt.Println(res)

	s = "abc"
	k = 3
	res = leetcode.TakeCharacters(s, k)
	fmt.Println(res)
}

// func main() {
// 	code := []int{5, 7, 1, 4}
// 	k := 2
// 	res := leetcode.MaximumSubarraySum(code, k)
// 	fmt.Println(res)

// 	code = []int{4, 4, 4}
// 	k = 3
// 	res = leetcode.MaximumSubarraySum(code, k)
// 	fmt.Println(res)
// }

// func main() {
// 	code := []int{5, 7, 1, 4}
// 	k := 3
// 	res := leetcode.Decrypt(code, k)
// 	fmt.Println(res)

// 	code = []int{2, 4, 9, 3}
// 	k = 2
// 	res = leetcode.Decrypt(code, k)
// 	fmt.Println(res)

// 	code = []int{2, 4, 9, 3}
// 	k = -2
// 	res = leetcode.Decrypt(code, k)
// 	fmt.Println(res)
// }

// func main() {
// 	var nums []int
// 	nums = []int{1,11,10}
// 	k := 3
// 	res := leetcode.ResultsArray(nums, k)
// 	fmt.Println(res)

// 	nums = []int{1,2,9,8}
// 	k = 2
// 	res = leetcode.ResultsArray(nums, k)
// 	fmt.Println(res)
// }
