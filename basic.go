package stats

import (
	//"fmt"
	"math"
)

func Min(sample []float64) float64 {
	min := math.MaxFloat64
	for i := range sample {
		if sample[i] < min {
			min = sample[i]
		}
	}
	return min
}

func Max(sample []float64) float64 {
	max := -1 * math.MaxFloat64
	for i := range sample {
		if sample[i] > max {
			max = sample[i]
		}
	}
	return max
}

func Sum(sample []float64) float64 {
	sum := float64(0)
	for i := range sample {
		sum += sample[i]
	}
	return sum
}

func Average(sample []float64) float64 {
	return Sum(sample) / float64(len(sample))
}
