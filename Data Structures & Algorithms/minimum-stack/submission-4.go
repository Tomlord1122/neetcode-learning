type MinStack struct {
	stk []int
	stkMin []int
}

func Constructor() MinStack {
	return MinStack{
		stk: []int{},
		stkMin: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stk = append(this.stk, val)
	if len(this.stkMin) != 0 && this.stkMin[len(this.stkMin)-1] < val{
		this.stkMin = append(this.stkMin, this.stkMin[len(this.stkMin)-1])
	} else {
		this.stkMin = append(this.stkMin, val)
	}
}

func (this *MinStack) Pop() {
	this.stk = this.stk[:len(this.stk)-1]
	this.stkMin = this.stkMin[:len(this.stkMin)-1]
}

func (this *MinStack) Top() int {
	return this.stk[len(this.stk)-1]
}

func (this *MinStack) GetMin() int {
	return this.stkMin[len(this.stkMin)-1]
}
