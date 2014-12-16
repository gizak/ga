package ga

import "errors"
import "math"

//import "fmt"

// Incubator contains all information needed in evolving
type Incubator struct {
	pool       chromPool
	cache      chromPool
	best       Chromosome
	iterCnt    int
	avgFitness float64
	config     Config
}

type chromPool []Chromosome

// NewIncubator returns a blank incubator, need run Init afterward
func NewIncubator(cfg Config) *Incubator {
	inc := Incubator{config: cfg}
	return &inc
}

// Init initialize the incubator
func (inc *Incubator) Init() error {
	// init pool
	pop := inc.config.Population
	chlen := inc.config.ChromLen
	chmax := inc.config.MaxChromEl
	inc.pool = make([]Chromosome, pop)

	// chk FitnessFunc
	if inc.config.FitnessFunc == nil {
		return errors.New("Need provide FitnessFunc")
	}

	// fill pool
	switch inc.config.EncodingMethod {
	case binary:
		for i := range inc.pool {
			inc.pool[i] = newBinChrom(chlen)
		}
	case intVec:
		for i := range inc.pool {
			inc.pool[i] = newIntChrom(chlen, chmax)
		}
	case orderedIntVec:
		for i := range inc.pool {
			inc.pool[i] = newOrderedIntChrom(chlen)
		}
	default:
		return errors.New("No matched encoding method")
	}

	// callback if not nil
	if fn := inc.config.InitCallback; fn != nil {
		fn()
	}
	return nil
}

func (inc *Incubator) doCrossover() {
	pop := inc.config.Population
	inc.cache = make([]Chromosome, 0, pop)

	switch inc.config.ParentSelMethod {
	case tour:
		for i := 0; i < pop/2; i++ {
			par1, par2 := tourSelPair(inc.pool, inc.config.TourSize, inc.config.TourRate)
			ic1 := par1.(*intChrom)
			ic2 := par2.(*intChrom)
			nsl1, nsl2 := onePointXover(ic1.value, ic2.value)

			nc1 := &intChrom{value: nsl1}
			nc2 := &intChrom{value: nsl2}
			inc.cache = append(inc.cache, nc1, nc2)
		}
	default:
		return
	}

	inc.cache, inc.pool = inc.pool, inc.cache
}

func (inc *Incubator) doSelection() {}

func (inc *Incubator) doMutation() {
	switch inc.config.MutationMethod {
	case flipBits:
		p := inc.config.MutationRate
		for i := range inc.pool {
			binFlipBits(inc.pool[i].Value().([]int), p)
		}
	default:
		return
	}
}

//
func (inc *Incubator) updateStatus() {
	bestFitness := math.Inf(-1)
	if inc.best != nil {
		bestFitness = inc.best.Fitness()
	}

	fn := inc.config.FitnessFunc
	sum := 0.0
	for i, v := range inc.pool {
		nfit := fn(v)
		sum += nfit
		inc.pool[i].SetFitness(nfit)
		if nfit > bestFitness {
			inc.best = inc.pool[i].Copy()
			bestFitness = nfit
		}
	}

	inc.avgFitness = sum / float64(len(inc.pool))
	inc.iterCnt++
}

func (inc *Incubator) AverageFitness() float64 {
	return inc.avgFitness
}

// Next returns iterable
func (inc *Incubator) Next() bool {
	if inc.iterCnt >= inc.config.MaxIterCnt {
		return false
	}

	inc.updateStatus()
	inc.doCrossover()
	inc.doMutation()
	inc.doSelection()

	return true
}
