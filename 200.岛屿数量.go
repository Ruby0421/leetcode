/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
    count:=0
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            if grid[i][j]=='1'{
                dfs(grid,i,j)
                count++
            }
        }
    }
    return count
}
func dfs(grid [][]byte,i int,j int){
    if i<0||j<0||i>=len(grid)||j>=len(grid[0])||grid[i][j]!='1'||grid[i][j]=='*'{
        return;
    }
    // temp:=grid[i][j]
    grid[i][j]='*'
    dfs(grid,i-1,j)
    dfs(grid,i+1,j)
    dfs(grid,i,j-1)
    dfs(grid,i,j+1)
    // grid[i][j]=temp
}

// @lc code=end

