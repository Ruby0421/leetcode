/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
    res:=[]int{}
    i:=0
    num2:=make([]int,len(nums))
    copy(num2,nums)
    num2=append(num2,nums...)
    // fmt.Println(num2)
    // fmt.Println(len(nums))
    for index,value:=range nums{
        position:=-1
        for i=index+1;i<len(num2);i++{
            if num2[i]>value{
                position=num2[i]
                break
            } 
        }
        res=append(res,position)
    }
    return res
}
// @lc code=end

