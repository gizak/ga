package ga

import "math/rand"

func binFlipBits(b []int, prob float64) {
	for i, v := range b {
		if rand.Float64() < prob {
			if v == 0 {
				b[i] = 1
			} else {
				b[i] = 0
			}
		}
	}
}
