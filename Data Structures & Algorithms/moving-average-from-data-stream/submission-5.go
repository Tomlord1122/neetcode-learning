type MovingAverage struct {
	size int
	window []int
}

func Constructor(size int) *MovingAverage {
	return &MovingAverage{
		size: size,
		window: []int{},
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.window = append(this.window, val)
	if len(this.window) > this.size{
		// pop the first element
		this.window = this.window[1:]
	}
	// Calculate the avg
	var avg float64
	sum := 0
	for _, val := range this.window{
		sum += val
	}
	avg = float64(sum) / float64(len(this.window))
	return avg
}
