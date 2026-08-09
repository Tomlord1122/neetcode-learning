type Node struct{
	isWord bool
	children map[rune]*Node
}

type PrefixTree struct {
	root *Node
}

func Constructor() PrefixTree {
    return PrefixTree{
		root: &Node{
			children: make(map[rune]*Node),
		},
	}
}

func (this *PrefixTree) Insert(word string) {
	cur := this.root
	for _, c := range word{
		if cur.children[c] == nil{
			cur.children[c] = &Node{children: make(map[rune]*Node)}
		}
		cur = cur.children[c]
	}
	cur.isWord = true
}

func (this *PrefixTree) Search(word string) bool {
	cur := this.root
	for _, c := range word{
		if cur.children[c] == nil{
			return false
		}
		cur = cur.children[c]
	}
	return cur.isWord
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


