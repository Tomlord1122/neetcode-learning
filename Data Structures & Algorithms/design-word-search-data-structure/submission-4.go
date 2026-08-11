type TrieNode struct{
	children map[rune]*TrieNode
	isWord bool
}
type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

func (this *WordDictionary) AddWord(word string)  {
	cur := this.root
	for _, c := range word{
		if cur.children[c] == nil{
			cur.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		cur = cur.children[c]
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
				if cur.children[rune(c)] == nil{
					return false
				}
				cur = cur.children[rune(c)]
			}
		}
		return cur.isWord
	}

	return dfs(0, this.root)
}
