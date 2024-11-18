package leetcode

func Decrypt(code []int, k int) []int {
	n := len(code)
	res := make([]int, n)
	if k == 0 {
		return res
	}
	if k > 0 {
		for i := 0; i < n; i++ {
			for j := 1; j <= k; j++ {
				res[i] += code[(i+j)%n]
			}
		}
	} else {
		for i := 0; i < n; i++ {
			for j := 1; j <= -k; j++ {
				res[i] += code[(i-j+n)%n]
			}
		}
	}
	return res
}