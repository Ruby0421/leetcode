/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	res:=[][]int{}
	sort.Slice(intervals,func(i,j int)bool{
		return intervals[i][0]<intervals[j][0]
	})
    if len(intervals)==0{
        return res
    }
	res=append(res,intervals[0])
	target:=intervals[0]
	for i:=0;i<len(intervals);i++{
		if target[1]<intervals[i][0]{
			res=append(res,intervals[i])
			target=intervals[i]
		}else if target[1]<intervals[i][1]{
			target[1]=intervals[i][1]
		}
	}
	return res
}
// @lc code=end

