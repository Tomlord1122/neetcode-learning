type MinStack struct {
    minStk []int
    stk []int
}

func Constructor() MinStack {
    return MinStack{
        minStk: []int{},
        stk: []int{},
    }
}

func (this *MinStack) Push(val int) {
    this.stk = append(this.stk, val)
    if len(this.minStk) != 0 && this.minStk[len(this.minStk)-1] <= val{
        this.minStk = append(this.minStk, this.minStk[len(this.minStk)-1])
    } else {
        this.minStk = append(this.minStk, val)
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


// Only getMin() is unique than the classical stack
// The naive way is that we can maintain another stack to keep track of the minimum value

// The overall complexity should be the samce (time and space)
