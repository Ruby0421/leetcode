/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
    res:=[]string{}
    left:=n
    right:=n
    generate(n,"",left,right,&res)
    return res
}
func generate (n int,s string,left int,right int, res *[]string){
    if left==0&&right==0{
		*res=append(*res,s)
        return ;
    }
    if left>0{
        generate(n,s+"(",left-1,right,res)
    }
    if left<right{
        generate(n,s+")",left,right-1,res)
    }
    return ;
}
// @lc code=end

