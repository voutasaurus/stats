package stats

import (
	//"fmt"
	"math"
)

func Variance(sample []float64, wholepop bool) float64 {
	//	sample := []float64{20,22,1,7,4}
	//	wholepop := false

	avg := Average(sample)

	// variance
	squarediff := float64(0)
	for i := range sample {
		squarediff += (sample[i] - avg) * (sample[i] - avg)
	}

	variance := float64(0)
	if wholepop {
		variance = squarediff / float64(len(sample))
	} else {
		variance = squarediff / float64(len(sample)-1)
	}

	/* // Debug:
	fmt.Println("avg:", avg)
	fmt.Println("variance:", variance)
	fmt.Println("standard deviation:", math.Sqrt(variance))
	*/

	return variance
}

func StdDev(sample []float64, wholepop bool) float64 {
	return math.Sqrt(Variance(sample, wholepop))
}
