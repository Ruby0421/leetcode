/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
    res:=[][]int{}
    temp:=[]int{}
    used:=make([]bool,len(nums))
    permuteCore(nums,0,temp,used,&res)
    return res
}
func permuteCore(nums []int,index int,temp []int,used []bool,res *[][]int){
    if len(temp)==len(nums){
        a:=make([]int,len(nums))
        copy(a,temp)
        *res=append(*res,a)
        return ;
    }
    for i:=0;i<len(nums);i++{
        // fmt.Println("nums:",nums)
        if i>index&&nums[i]==nums[i-1]{
            continue ;
        }
        if used[i]!=true{
            used[i]=true
            temp=append(temp,nums[i])
            // fmt.Println("1temp:",temp)
            permuteCore(nums,i+1,temp,used,res)
            temp=temp[:len(temp)-1]
            used[i]=false
        }
        // fmt.Println("2temp:",temp)
    }
    return ;
}
// @lc code=end

