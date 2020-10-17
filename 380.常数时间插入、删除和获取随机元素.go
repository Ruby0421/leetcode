/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] 常数时间插入、删除和获取随机元素
 */

// @lc code=start
type RandomizedSet struct {
    number []int
    // isHave map[int]bool
    position map[int]int
    Len int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    // m:=RandomizedSet{
    //     number: []int{},
    //     isHave: map[int]bool{},
    // }
    m:=RandomizedSet{
        number: []int{},
        position: map[int]int{},
        Len:0,
    }
    return m
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    // if this.isHave[val]==false{
    //     this.number=append(this.number,val)
    //     this.isHave[val]=true
    //     return true
    // }
    _,ok:=this.position[val]
    if !ok{
        this.number=append(this.number,val)
        // fmt.Println(this.number)
        this.position[val]=len(this.number)-1
        this.Len=this.Len+1
        // fmt.Println("insert")
        // fmt.Println(this)
        return true
    }
    return false
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    // res:=[]int{}
    // if this.isHave[val]==true{
    //     for i:=0;i<len(this.number);i++{
    //         if this.number[i]!=val{
    //             res=append(res,this.number[i])
    //         }
    //     }
    //     return true
    // }
    // return false
    _,ok:=this.position[val]
    if ok{
            m:=this.number[this.Len-1]
            index:=this.position[val]
            this.number[index]=m
            this.number=this.number[:this.Len-1]
            this.position[m]=index
            this.Len=this.Len-1
            delete (this.position,val)
            return true
        }
    // for k,v := range this.position{
    //     fmt.Println("k:%v,v:%v",k,v)
    // }
  
    return false
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    // fmt.Println("random")
    // fmt.Println(this)
    if len(this.number)!=0{
        rand.Seed(time.Now().UnixNano())
        num := rand.Intn(len(this.number))
        return this.number[num]
    }
    return -1
}
/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

