/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
    right:=0
    for i:=0;i<len(nums);i++{
        if i>right{
            return false
        }
        dis:=nums[i]+i
        right=max(dis,right)
        if right>=len(nums)-1{
            return true
        }
    }
    return false
}
func max(a,b int)int{
    if a>b{
        return a
    }else{
        return b
    }
}
// @lc code=end

