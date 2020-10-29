/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] 最大连续1的个数 III
 */

// @lc code=start
func longestOnes(A []int, K int) int {
	left:=0
	nums0:=0
	maxNum:=0
	for i:=0;i<len(A);i++{
		if A[i]==0{
			nums0++
		}
		for nums0>K{
			if A[left]==0{
				nums0--
			}
			left++
		}
		maxNum=max(maxNum,i-left+1)
	}
	return maxNum
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
// @lc code=end

