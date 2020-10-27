/*
 * @lc app=leetcode.cn id=914 lang=golang
 *
 * [914] 卡牌分组
 */

// @lc code=start
func hasGroupsSizeX(deck []int) bool {
    m:=map[int]int{}
    for i:=0;i<len(deck);i++{
        m[deck[i]]++
    }
    maxValue:=m[deck[0]]
    for _,v :=range m{
        maxValue=gcd(v,maxValue)
        if maxValue<2{
            return false
        }
    }
    return true
}

func gcd( a int,b int)int{
    temp:=a%b
    if temp>0{
        return gcd(b,temp)
    }else{
        return b
    }
}
// @lc code=end

