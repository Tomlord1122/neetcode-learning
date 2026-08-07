type MinStack struct {
	stk []int
	minStk []int
}

func Constructor() MinStack {
	return MinStack{
		stk: []int{},
		minStk: []int{},
	}
}

func (this *MinStack) Push(val int) {
	// push stk and track the min in minStk
	this.stk = append(this.stk, val)
	if len(this.minStk) == 0{
		this.minStk = append(this.minStk, val)
	} else {
		this.minStk = append(this.minStk, min(this.minStk[len(this.minStk)-1], val))
	}
}

func (this *MinStack) Pop() {
	this.stk = this.stk[:len(this.stk)-1]
	this.minStk = this.minStk[:len(this.minStk)-1]
}

func (this *MinStack) Top() int {
	return this.stk[len(this.stk)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStk[len(this.minStk)-1]
}
