package discreterand

import (
	"math"
	"math/rand"
	"testing"
)

func TestAlias(t *testing.T) {

	p := []float64{0.125, 0.2, 0.1, 0.25, 0.1, 0.1, 0.125}
	counts := make([]int, len(p))

	a := NewAlias(p, rand.NewSource(0))

	const rounds = 10e6

	for i := 0; i < rounds; i++ {
		n := a.Next()
		counts[n]++
	}

	const eps = 0.001
	for i := range p {
		g := float64(counts[i]) / float64(rounds)
		if math.Abs(g-p[i]) > eps {
			t.Errorf("failed alias method test element %d: got %f expected %f +/- %f", i, g, p[i], eps)
		}
	}
}
