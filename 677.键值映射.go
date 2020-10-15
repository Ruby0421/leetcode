/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 */

// @lc code=start
type MapSum struct {
    node map[string]int
}

/** Initialize your data structure here. */
func Constructor() MapSum {
    m:=MapSum{}
    m.node=map[string]int{}
    return m
}


func (this *MapSum) Insert(key string, val int)  {
    this.node[key]=val
}


func (this *MapSum) Sum(prefix string) int {
    count:=0
    for key,value :=range this.node{
        if len(prefix)<=len(key)&&key[:len(prefix)]==prefix{
            count=count+value
        }
    }
    return count
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end

