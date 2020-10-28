/*
 * @lc app=leetcode.cn id=1109 lang=golang
 *
 * [1109] 航班预订统计
 */

// @lc code=start
// func corpFlightBookings(bookings [][]int, n int) []int {
//     res:=[]int{}
//     m:=map[int]int{}
//     for i:=0;i<len(bookings);i++{
//         first:=bookings[i][0]
//         last:=bookings[i][1]
//         value:=bookings[i][2]
//         for j:=first;j<=last;j++{
//             m[j]=m[j]+value
//         }
//     }
//     for i:=0;i<n;i++{
//         res=append(res,m[i+1])
//     }
//     return res
// }
// func corpFlightBookings(bookings [][]int, n int) []int {
//     res:=make([]int,n)
//     for i:=0;i<len(bookings);i++{
//         for j:=bookings[i][0];j<=bookings[i][1];j++{
//             res[j-1]=res[j-1]+bookings[i][2]
//         }
//     }
//     return res
// }
func corpFlightBookings(bookings [][]int, n int) []int {
    res:=make([]int,n)
    for i:=0;i<len(bookings);i++{
        res[bookings[i][0]-1]=res[bookings[i][0]-1]+bookings[i][2]
        if bookings[i][1]<n{
            res[bookings[i][1]]=res[bookings[i][1]]-bookings[i][2]
        }
    }
    for i:=1;i<n;i++{
        res[i]=res[i-1]+res[i]
    }
    return res
}
// @lc code=end

