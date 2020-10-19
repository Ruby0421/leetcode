/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
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
    for i:=start;i<len(candidates);i++ {
        if candidates[i]>target{
            return ;
        }
        if i>start&&candidates[i]==candidates[i-1]{
            continue
        }
        temp=append(temp,candidates[i])
        combine(candidates,i+1,target-candidates[i],temp,res)
        temp=temp[:len(temp)-1]
    }
    return ;
}
// @lc code=end

