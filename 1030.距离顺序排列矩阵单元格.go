/*
 * @lc app=leetcode.cn id=1030 lang=golang
 *
 * [1030] 距离顺序排列矩阵单元格
 */

// @lc code=start
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
    hashMap:=map[int][][]int{}
    minValue:=abs(0,r0)+abs(0,c0)
    maxValue:=abs(0,r0)+abs(0,c0)
    for i:=0;i<R;i++{
        for j:=0;j<C;j++{
            dis:=abs(i,r0)+abs(j,c0)
            hashMap[dis]=append(hashMap[dis],[]int{i,j})
            minValue=min(minValue,dis)
            maxValue=max(maxValue,dis)
        }
    }
    res:=[][]int{}
    for i:=minValue;i<=maxValue;i++{
        res=append(res,hashMap[i]...)
    }
    return res
}
func abs(a,b int)int{
    if a>b{
        return a-b
    }
    return b-a
}
func min(a,b int)int{
    if a<b{
        return a
    }
    return b
}
func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
// @lc code=end

