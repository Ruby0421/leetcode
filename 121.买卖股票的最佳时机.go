/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
    maxPrice:=0
    for i:=0;i<len(prices)-1;i++{
        for j:=i+1;j<len(prices);j++{
            if prices[j]-prices[i]>0{
                maxPrice=max(maxPrice,prices[j]-prices[i])
            }
        }
    }
    return maxPrice
}
func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
// @lc code=end

