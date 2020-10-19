/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
    res:=[][]int{}
    temp:=[]int{}
    sort.Ints(nums)
    used:=make([]bool,len(nums))
    permuteUniqueCore(nums,temp,used,&res)
    return res
}
func permuteUniqueCore(nums []int,temp []int,used []bool,res *[][]int){
    if len(temp)==len(nums){
        a:=make([]int,len(nums))
        copy(a,temp)
        *res=append(*res,a)
    }
    for i:=0;i<len(nums);i++{
        if i>0&&nums[i]==nums[i-1]&&used[i-1]{
            continue
        }
        if !used[i]{
            used[i]=true
            temp=append(temp,nums[i])
            // fmt.Println("temp:",temp)
            permuteUniqueCore(nums,temp,used,res)
            temp=temp[:len(temp)-1]
            used[i]=false
        }
    }
    return ;
}
// @lc code=end

