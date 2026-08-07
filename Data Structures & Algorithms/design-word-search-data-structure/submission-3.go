type TrieNode struct{
	children [26]*TrieNode
	isWord bool
}
type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{
		root: &TrieNode{},
	}
}

func (this *WordDictionary) AddWord(word string)  {
	cur := this.root
	for _, c := range word{
		index := c - 'a'
		if cur.children[index] == nil{
			cur.children[index] = &TrieNode{}
		}
		cur = cur.children[index]
	}
	cur.isWord = true
}

func (this *WordDictionary) Search(word string) bool {

	var dfs func(i int, node *TrieNode) bool
	dfs = func(i int, node *TrieNode) bool{
		cur := node
		for j := i; j < len(word); j++{
			c := word[j]
			if c == '.'{
				for _, child := range cur.children{
					if child != nil && dfs(j+1, child){
						return true
					}
				}
				return false
			} else {
				index := c - 'a'
				if cur.children[index] == nil{
					return false
				}
				cur = cur.children[index]
			}
		}
		return cur.isWord
	}
	return dfs(0, this.root)
}
