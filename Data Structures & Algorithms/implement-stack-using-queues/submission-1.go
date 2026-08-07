type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		queue1: []int{},
		queue2: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.queue1 = append(this.queue1, x)
}

func (this *MyStack) Pop() int {
	// pop out all element in queue1 
	res := 0
	for len(this.queue1) > 0{
		if len(this.queue1) == 1{
			res = this.queue1[0]
			this.queue1 = this.queue1[1:]
			break
		}
		val := this.queue1[0]
		this.queue1 = this.queue1[1:]
		this.queue2 = append(this.queue2, val)
	}

	// Then revert back
	for len(this.queue2) > 0{
		val := this.queue2[0]
		this.queue2 = this.queue2[1:]
		this.queue1 = append(this.queue1, val)
	}
	return res
}

func (this *MyStack) Top() int {
	return this.queue1[len(this.queue1)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */



