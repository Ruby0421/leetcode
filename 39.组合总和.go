/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
    res:=[][]int{}
    temp:=[]int{}
    sort.Ints(candidates)
    combine(candidates,0,target,temp,&res)
    return res
}
func combine(candidates []int,start int,target int,temp []int,res *[][]int){
    if target==0{
        a:=make([]int,len(temp))
        copy(a,temp)
        *res=append(*res,a)
        return ;
    }
    for i:=start;i<len(candidates);i++{
        if candidates[i]>target{
            break ;
        }
        temp=append(temp,candidates[i])
        combine(candidates,i,target-candidates[i],temp,res)
        temp=temp[:len(temp)-1]
    }
    return ;
}
// @lc code=end

