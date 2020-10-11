/*
 * @lc app=leetcode.cn id=1071 lang=golang
 *
 * [1071] 字符串的最大公因子
 */

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {
    if str1+str2!=str2+str1{
        return ""
    }
    if len(str1)>len(str2){
        return str1[:gcd(len(str1),len(str2))]
    }else{
        return str2[:gcd(len(str2),len(str1))]
    }
}
func gcd(a,b int)int{
    if b==0{
        return a
    }else{
        return gcd(b,a%b)
    }

}
// @lc code=end

