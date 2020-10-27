/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	for i:=0;i<len(nums)-1;i++{
		sum:=nums[i]
		for j:=i+1;j<len(nums);j++{
			sum=sum+nums[j]
			if sum==k||k!=0&&sum%k==0{
				return true
			}
		}
	}
	return false
}
// @lc code=end

