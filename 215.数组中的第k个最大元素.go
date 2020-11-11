/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
    sort.Slice(nums,func(i,j int)bool{
        return nums[i]>=nums[j]
    })
    return nums[k-1]
}
// @lc code=end

