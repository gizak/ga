package ga

import "sort"
import "math/rand"

func (cp chromPool) Len() int {
	return len(cp)
}

func (cp chromPool) Swap(i, j int) {
	cp[i], cp[j] = cp[j], cp[i]
}

func (cp chromPool) Less(i, j int) bool {
	return cp[i].Fitness() < cp[j].Fitness()
}

func tourSelOne(pool chromPool, size int, p float64) Chromosome {
	chs := chromPool(make([]Chromosome, size))
	l := len(pool)
	for i := range chs {
		chs[i] = pool[rand.Intn(l)]
	}

	sort.Sort(chs)

	if rand.Float64() > p {
		return chs[size-2]
	}
	return chs[size-1]
}

func tourSelPair(pool chromPool, size int, p float64) (Chromosome, Chromosome) {
	return tourSelOne(pool, size, p), tourSelOne(pool, size, p)
}
