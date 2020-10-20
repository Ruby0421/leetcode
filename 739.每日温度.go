/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
    res:=[]int{}
    j:=0
    for i:=0;i<len(T);i++{
        for j=i+1;j<len(T);j++{
            if T[j]>T[i]{
                res=append(res,j-i)
                break
            }
        }
        if i!=len(T)-1&&j==len(T){
            res=append(res,0)
        }
    }
    res=append(res,0)
    return res
}
// @lc code=end

