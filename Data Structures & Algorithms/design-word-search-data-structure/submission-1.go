type TrieNode struct{
	Children [26]*TrieNode
	Word bool
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
		index := c -'a'
		if cur.Children[index] == nil{
			cur.Children[index] = &TrieNode{}
		}
		cur = cur.Children[index]
	}
	cur.Word = true
}

func (this *WordDictionary) Search(word string) bool {
    var dfs func(i int, node *TrieNode) bool
	dfs = func(i int, node *TrieNode) bool{
		cur := node
		for j := i; j < len(word); j++{
			c := word[j]
			if c == '.'{
				for _, child := range cur.Children{
					if child != nil && dfs(j+1, child){
						return true
					}
				}
				return false
			} else {
				index := c - 'a'
				if cur.Children[index] == nil{
					return false 
				}
				cur = cur.Children[index]
			}
		}
		return cur.Word
	}

	return dfs(0, this.root)
}
