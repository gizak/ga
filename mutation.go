package ga

import "math/rand"

func (bc *BitStrChromosome) Mutate(p float64) {
	for i, v := range bc.Value() {
		if rand.Float64() < p {
			if v == 0 {
				bc.value[i] = 1
			} else {
				bc.value[i] = 0
			}
		}
	}
}

func (vc *IntVecChromosome) Mutate(p float64, min, max int) {
	for i := range vc.Value() {
		if rand.Float64() < p {
			vc.value[i] = rand.Intn(max-min+1) + min
		}
	}
}
