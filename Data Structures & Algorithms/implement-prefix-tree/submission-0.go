type TrieNode struct{
	children map[rune]*TrieNode
	endOfWord bool
}

type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
	return PrefixTree{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (this *PrefixTree) Insert(word string) {
	cur := this.root
	for _, c := range word{
		if cur.children[c] == nil{
			cur.children[c] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		cur = cur.children[c]
	}
	cur.endOfWord = true

}

func (this *PrefixTree) Search(word string) bool {
	cur := this.root
	for _, c := range word{
		if cur.children[c] == nil{
			return false
		}
		cur = cur.children[c]
	}
	return cur.endOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	cur := this.root
	for _, c := range prefix{
		if cur.children[c] == nil{
			return false
		}
		cur = cur.children[c]
	}
	return true
}
