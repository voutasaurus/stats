package stats

import (
	"fmt"
	"testing"
)

func TestVariance(t *testing.T) {
	sample := []float64{20,22,1,7,4}
	wholepop := false

	v := Variance(sample, wholepop)
	
	fmt.Println("Test variance: ", v)
	
	return
}