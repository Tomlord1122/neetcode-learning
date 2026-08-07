type MinStack struct {
	stk1 []int
	stk2 []int
}

func Constructor() MinStack {
	return MinStack{
		stk1: []int{},
		stk2: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stk1 = append(this.stk1, val)
	if len(this.stk2) == 0{
		this.stk2 = append(this.stk2, val)
	} else {
		this.stk2 = append(this.stk2, min(val, this.stk2[len(this.stk2)-1]))
	}
}

func (this *MinStack) Pop() {
	this.stk1 = this.stk1[:len(this.stk1)-1]
	this.stk2 = this.stk2[:len(this.stk2)-1]
}

func (this *MinStack) Top() int {
	return this.stk1[len(this.stk1)-1]
}

func (this *MinStack) GetMin() int {
	return this.stk2[len(this.stk2)-1]
}
