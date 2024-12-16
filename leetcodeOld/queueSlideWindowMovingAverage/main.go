package main

type MovingAverage struct {
	list []int
	sum  int
	size int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.list) == this.size {
		this.sum -= this.list[0]
		this.list = this.list[1:]
	}
	this.list = append(this.list, val)
	this.sum += val
	return float64(this.sum) / float64(len(this.list))
}
