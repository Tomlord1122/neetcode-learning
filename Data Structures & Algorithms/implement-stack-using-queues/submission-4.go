type MyStack struct {
	q1, q2 []int
}

func Constructor() MyStack {
	return MyStack{
		q1: []int{},
		q2: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.q1 = append(this.q1, x)
}

func (this *MyStack) Pop() int {
	// queue has FIFO property
	for len(this.q1) > 1 {
		cur := this.q1[0]
		this.q1 = this.q1[1:]
		this.q2 = append(this.q2, cur)
	}
	// pop the last element
	res := this.q1[0]
	this.q1 = this.q1[1:]
	// push it back to q1
	for len(this.q2) > 0 {
		cur := this.q2[0]
		this.q2 = this.q2[1:]
		this.q1 = append(this.q1, cur)
	}
	return res
}

func (this *MyStack) Top() int {
	res := this.Pop()
	this.Push(res)
	return res
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
