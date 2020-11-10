/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int)  {
    left:=-1
    target:=nums[len(nums)-1]
    for i:=len(nums)-2;i>=0;i--{
        if nums[i]<target{
            left=i
            break
        }else{
            target=nums[i]
        }
    }
    if left!=-1{
        index:=left+1
        for i:=left+1;i<len(nums);i++{
            if nums[i]>nums[left]{
                if nums[i]<=nums[index]{
                    index=i
                }
            }
        }
        nums[left],nums[index]=nums[index],nums[left]
        sort.Ints(nums[left+1:])
    }else{
        sort.Ints(nums)
    }
}
// @lc code=end

