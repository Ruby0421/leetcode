/*
 * @lc app=leetcode.cn id=1208 lang=golang
 *
 * [1208] 尽可能使字符串相等
 */

// @lc code=start
func equalSubstring(s string, t string, maxCost int) int {
    count:=0
    i,j:=0,0
    sum:=0
    for i<len(s){
        sum=sum+absMinus(int(s[i]),int(t[i]))
        i++
        for sum>maxCost{
            sum=sum-absMinus(int(s[j]),int(t[j]))
            j++
        }
        count=max(count,i-j)
    }
    return count
}
func absMinus(a int,b int)int{
    if a>b{
        return a-b
    }
    return b-a
}
func max(a int,b int)int{
    if a>b{
        return a
    }
    return b
}
// @lc code=end

