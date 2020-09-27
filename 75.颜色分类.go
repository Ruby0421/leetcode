/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int)  {
	left,right:=0,len(nums)-1
    i:=0
    for i<=right{
        if nums[i]==0{
            nums[left],nums[i]=nums[i],nums[left]
            left++
            i++
        }else if nums[i]==1{
            i++
        }else{
            nums[i],nums[right]=nums[right],nums[i]
            right--
        }
    }
}
// @lc code=end

