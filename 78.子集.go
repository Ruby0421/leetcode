/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
var res [][]int
func subsets(nums []int) [][]int {
	res=make([][]int,0)
	temp:=[]int{}
	dfs(nums,0,temp)
	return res
}
func dfs(nums []int,index int,temp []int) {
	if len(nums)==index{
		cur :=make([]int,len(temp))
		copy(cur,temp)
		res=append(res,cur)
		return 
	}
	dfs(nums,index+1,temp)
	temp=append(temp,nums[index])
	dfs(nums,index+1,temp)
	temp=temp[:len(temp)-1]
}
// @lc code=end

