/*
 * @lc app=leetcode.cn id=1486 lang=golang
 *
 * [1486] 数组异或操作
 */

// @lc code=start
func xorOperation(n int, start int) int {
    nums:=make([]int,n)
    res:=0
    for i:=0;i<n;i++{
        nums[i]=start+2*i
        res=res^nums[i]
    }
    return res
}
// @lc code=end

