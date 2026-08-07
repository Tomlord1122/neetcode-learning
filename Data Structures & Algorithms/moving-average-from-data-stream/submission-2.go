type MovingAverage struct {
	queue []int
	size int
}

func Constructor(size int) *MovingAverage {
	return &MovingAverage{
		queue: []int{},
		size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	// add the val
	this.queue = append(this.queue, val)
	// check the length condition
	if len(this.queue) > this.size{
		// move the window right by 1
		this.queue = this.queue[1:]
	}
	// caluclate the average and return it
	return float64(sum(this.queue)) / float64(len(this.queue))
}

func sum(queue []int) int{
	res := 0
	for _, val := range queue{
		res += val
	}
	return res
}
