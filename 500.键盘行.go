/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 */

// @lc code=start
func findWords(words []string) []string {
    res:=[]string{}
    a :=[]byte{'Q','W','E','R','T','Y','U','I','O','P'} 
    b :=[]byte{'A','S','D','F','G','H','J','K','L'}
    c :=[]byte{'Z','X','C','V','B','N','M'}  
    m:=make(map[byte]int)
    for i:=0;i<len(a);i++{
        m[a[i]]=1
        m[a[i]+32]=1
    }
    for i:=0;i<len(b);i++{
        m[b[i]]=2
        m[b[i]+32]=2
    }
    for i:=0;i<len(c);i++{
        m[c[i]]=3
        m[c[i]+32]=3
    }
    for i:=0;i<len(words);i++{
        x:=[]byte(words[i])
        count:=0
        for j:=1;j<len(x);j++{
            if m[x[j]]==m[x[j-1]]{
                count++
            }
        }
        if count==len(x)-1{
            res=append(res,words[i])
        }
    }
    return res  
}
// @lc code=end

