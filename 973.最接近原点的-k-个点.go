/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */

// @lc code=start
// func kClosest(points [][]int, K int) [][]int {
//     res:=[][]int{}
//     sortNums:=[]int{}
//     hashMap:=map[int][]int{}
//     for i:=0;i<len(points);i++{
//         sum:=points[i][0]*points[i][0]+points[i][1]*points[i][1]
//         sortNums=append(sortNums,sum)
//         _,ok:=hashMap[sum]
//         if !ok{
//             hashMap[sum]=[]int{i}
//         }else{
//             hashMap[sum]=append(hashMap[sum],i)
//         }
//     }
//     sort.Ints(sortNums)
//     m:=0
//     for i:=0;m<K;i++{
//         for j:=0;j<len(hashMap[sortNums[i]]);j++{
//             res=append(res,points[hashMap[sortNums[i]][j]])
//             m++
//         }
//     }
//     return res
// }

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points,func (i,j int)bool{
		return points[i][0]*points[i][0]+points[i][1]*points[i][1]<points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})
	return points[:K]
}
// @lc code=end

