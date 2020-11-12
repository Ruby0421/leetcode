/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    newList:=&ListNode{
        Val:0,Next:nil,
    }
    root:=newList
    index:=0
    for l1!=nil||l2!=nil||index>0{
        subList:=&ListNode{
            Val:0,Next:nil,
        }
        root.Next=subList
        root=subList
        if l1!=nil{
            index=index+l1.Val
            l1=l1.Next
        }
        if l2!=nil{
            index=index+l2.Val
            l2=l2.Next
        }
        fmt.Println(index)
        a:=index%10
        root.Val=a
        // fmt.Println(root.Val)
        index=index/10
        // fmt.Println(index)
    }
    return newList.Next
}
// @lc code=end

