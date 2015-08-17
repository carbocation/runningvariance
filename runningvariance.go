/*
runningvariance computes accurate running mean, variance, and standard deviation

It is based on code by John D Cook: http://www.johndcook.com/blog/standard_deviation/
*/
package runningvariance

import (
	"math"
)

type RunningStat struct {
	n    uint
	newM float64
	oldM float64
	newS float64
	oldS float64
}

func NewRunningStat() *RunningStat {
	return &RunningStat{}
}

func (r *RunningStat) Push(x float64) {
	r.n++

	// See Knuth TAOCP vol 2, 3rd edition, page 232
	if r.n == 1 {
		r.oldM = x
		r.newM = x
		r.oldS = 0.0
	} else {
		r.newM = r.oldM + (x-r.oldM)/float64(r.n)
		r.newS = r.oldS + (x-r.oldM)*(x-r.newM)

		// set up for next iteration
		r.oldM = r.newM
		r.oldS = r.newS
	}
}

func (r *RunningStat) NumDataValues() uint {
	return r.n
}

func (r *RunningStat) Mean() float64 {
	return r.newM
}

func (r *RunningStat) Variance() float64 {
	if r.n > 1 {
		return r.newS / (float64(r.n) - 1)
	}

	return 0.0
}

func (r *RunningStat) StandardDeviation() float64 {
	return math.Sqrt(r.Variance())
}
