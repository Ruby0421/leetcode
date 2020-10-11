/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
    // for i:=0;i<n;i++{
    //     nums1[m+i]=nums2[i]
    // }
    // sort.Ints(nums1) 
    
    count:=m+n-1
    index1:=m-1
    index2:=n-1
    for index1>=0&&index2>=0{
        nums1[count]=max(nums1[index1],nums2[index2])
        count--
        if nums1[index1]>nums2[index2]{
            index1--
        }else{
            index2--
        }
    }
    for index2>=0{
        nums1[count]=nums2[index2]
        count--
        index2--
    }
}
func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
// @lc code=end

