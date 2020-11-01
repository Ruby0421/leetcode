/*
 * @lc app=leetcode.cn id=617 lang=golang
 *
 * [617] 合并二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    root1:=t1
    root2:=t2
    root:=&TreeNode{Left:nil,Right:nil}
    if root1==nil{
        root=root2
        return root2
    }
    if root2==nil{
        root=root1
        return root1
    }        
    root.Val=root1.Val+root2.Val
    root.Left=mergeTrees(root1.Left,root2.Left)
    root.Right=mergeTrees(root1.Right,root2.Right)
    return root
}
// @lc code=end

