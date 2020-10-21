/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 朋友圈
 */

// @lc code=start
func findCircleNum(M [][]int) int {
    count:=0
    for i:=0;i<len(M);i++{
        for j:=0;j<len(M[0]);j++{
            if M[i][j]==1{
                count=count+1
                dfs(M,i)
                fmt.Println("count:",count)
            }
        }
    }
    return count
}
func dfs(M [][]int,i int){
    for j:=0;j<len(M[i]);j++{
        if M[i][j]==1{
            M[i][j]=0
            dfs(M,j)
        }
    }
}
// @lc code=end

