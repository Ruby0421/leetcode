/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
    sum:=0
    minLen:=len(nums)
    Len:=0
    for i:=0;i<len(nums);i++{
        sum=sum+nums[i]
    }
    if sum<s{
        return 0
    }
    for i:=0;i<len(nums);i++{
        sum:=0
        for j:=i;j<len(nums);j++{
            sum=sum+nums[j]
            // fmt.Println("sum:",sum)
            if sum>=s{
                Len=j-i+1
                fmt.Println("Len:",Len)
                break
            }
        }
        minLen=min(Len,minLen)
    }
    return minLen
}
func min(a int,b int)int{
    if a<b{
        return a
    }
    return b
}
// @lc code=end

