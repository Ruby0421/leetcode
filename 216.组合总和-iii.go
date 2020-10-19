/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
    nums:=[]int{1,2,3,4,5,6,7,8,9}
    res:=[][]int{}
    temp:=[]int{}
    combine(nums,0,k,n,temp,&res)
    return res
}
func combine(nums []int,index int,k int,n int,temp []int,res *[][]int) {
    if len(temp)==k&&n==0{
        a:=make([]int,len(temp))
        copy(a,temp)
        *res=append(*res,a)
        fmt.Println("res:",res)
        return ;
    }
    for i:=index;i<len(nums);i++{
        fmt.Println("index,i:",index,i)
        if nums[i]>n{
            return ;
        }
        if i>index&&nums[i]==nums[i-1]{
            continue;
        }
        temp=append(temp,nums[i])
        fmt.Println("temp:",temp)
        combine(nums,i+1,k,n-nums[i],temp,res)
        temp=temp[:len(temp)-1]
    }
    return ;
}
// @lc code=end

