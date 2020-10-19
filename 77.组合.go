/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
    nums:=make([]int,n)
    for i:=0;i<=n-1;i++{
        nums[i]=i+1
    }
    res:=[][]int{}
    temp:=[]int{}
    coreCombine(nums,0,k,temp,&res)
    return res
}
func coreCombine(nums []int,index int,k int,temp []int,res *[][]int){
    if len(temp)==k{
        a:=make([]int,k)
        copy(a,temp)
        *res=append(*res,a)
    }
    for i:=index;i<len(nums);i++{
        if i>index&&nums[i]==nums[i-1]{
            continue;
        }
        temp=append(temp,nums[i])
        coreCombine(nums,i+1,k,temp,res)
        temp=temp[:len(temp)-1]
    }
    return ;
}
// @lc code=end

