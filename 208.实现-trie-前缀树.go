/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	children map[byte]*Trie
	isWords bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    trie:=Trie{}
    trie.children=map[byte]*Trie{}
    trie.isWords=false
	// return Trie{children:map[string]Trie{},isEnd:true,isWords:false}
    return trie

}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    node:=this
    for i:=0;i<len(word);i++{
        _,ok:=node.children[word[i]]
        if !ok{
            node.children[word[i]]=&Trie{children:map[byte]*Trie{},isWords:false}
        }
        node=node.children[word[i]]
    }
    node.isWords=true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    node:=this
    for i:=0;i<len(word);i++{
        _,ok:=node.children[word[i]]
        if  !ok{
            return false
        }
        node=node.children[word[i]]
    }
    if node.isWords==true{
        return true
    }
    return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    node:=this
    for i:=0;i<len(prefix);i++{
        _,ok:=node.children[prefix[i]]
        if  !ok{
            return false
        }
        node=node.children[prefix[i]]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end




/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

