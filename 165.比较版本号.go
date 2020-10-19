/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
    str1:=strings.Split(version1,".")
    str2:=strings.Split(version2,".")
    maxlen:=max(len(str1),len(str2))
    a:=maxlen-len(str1)
    b:=maxlen-len(str2)
    if len(str1)<maxlen{
        for i:=0;i<a;i++{
            str1=append(str1,"0")
        }
    }
    if len(str2)<maxlen{
        for i:=0;i<b;i++{
            str2=append(str2,"0")
        }
    }
    fmt.Println(maxlen)
    for i:=0;i<maxlen;i++{
        value1,_:=strconv.Atoi(str1[i])
        value2,_:=strconv.Atoi(str2[i])
        if value1>value2{
            return 1
        }else if value1<value2{
            return -1
        }else{
            continue
        }
    }
    return 0;
}
func max(a int,b int)int{
    if a>b{
        return a
    }
    return b
}

// @lc code=end

