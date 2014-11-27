package ga

import "math/rand"
import "sort"

//import "fmt"

type bunchChromFitTester []IntVecChromUber

func (bc bunchChromFitTester) Len() int {
	return len(bc)
}

func (bc bunchChromFitTester) Less(i, j int) bool {
	return bc[i].Fitness() > bc[j].Fitness()
}

func (bc bunchChromFitTester) Swap(i, j int) {
	bc[i], bc[j] = bc[j], bc[i]
}

func tourSelectOne(ch []IntVecChromUber, size int, p float64) IntVecChromUber {
	//selected ch
	sch := bunchChromFitTester(make([]IntVecChromUber, size))
	n := len(ch)
	for i := 0; i < size; i++ {
		sch[i] = ch[rand.Intn(n)]
	}
	sort.Sort(sch)
	//fmt.Print(sch)
	if rand.Float64() > p {
		return sch[1]
	}
	return sch[0]
}

func tourSelectPair(ch []IntVecChromUber, size int, p float64) (IntVecChromUber, IntVecChromUber) {
	return tourSelectOne(ch, size, p), tourSelectOne(ch, size, p)
}
