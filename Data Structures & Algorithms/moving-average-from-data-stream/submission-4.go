type MovingAverage struct {
	window []int
	size int
}

func Constructor(size int) *MovingAverage {
	return &MovingAverage{
		window: []int{},
		size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	// append val
	this.window = append(this.window, val)
	// check size condition
	if len(this.window) > this.size{
		this.window = this.window[1:]
	}
	// retrun average
	sum := 0
	for _, num := range this.window{
		sum += num
	}
	return float64(sum) / float64(len(this.window))
}
