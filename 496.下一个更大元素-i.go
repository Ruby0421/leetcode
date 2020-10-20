/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    m:=map[int]int{}
    res:=[]int{}
    for i:=0;i<len(nums2);i++{
        m[nums2[i]]=i
    }
    j:=0
    for _,value:=range nums1{
        position:=-1
        for j=m[value];j<len(nums2);j++{
            if nums2[j]>value{
                position=nums2[j]
                break
            }
        }
        res=append(res,position)
    }
    return res
}
// @lc code=end

