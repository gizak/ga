package ga

import "errors"
import "math"

type Incubator interface {
	Init() error
	IsTerminal() bool
	UpdateFitness()
	UpdateCount()
	DoCrossover() error
	DoMutation() error
	DoSelection() error
	InitCb()
	EachCb()
	FinalCb()
}

func EvolveOnce(incu Incubator) error {
	if err := incu.Init(); err == nil {
		incu.InitCb()
		for !incu.IsTerminal() {
			incu.DoCrossover()
			incu.DoMutation()
			incu.UpdateFitness()
			incu.DoSelection()
			incu.UpdateCount()
			incu.EachCb()
		}
		incu.FinalCb()
	} else {
		return err
	}
	return nil
}

type IntVecIncubator struct {
	FitnessFunc     func([]int) float64
	Pool            []IntVecChromUber
	orldPool        []IntVecChromUber
	InitCallback    func()
	EachCallback    func()
	FinalCallback   func()
	MaxIterCnt      int
	Population      int
	MutationRate    float64
	TourRate        float64
	TourSize        int
	MinChromEl      int
	MaxChromEl      int
	ChromLen        int
	Best            IntVecChromUber
	BestFitness     float64
	IterCnt         int
	CrossoverRate   float64
	SelectionMethod string
	EncodingMethod  string
	CrossoverMethod string
	MutationMethod  string
	ParentSelMethod string
}

func NewIntVecIncubator() *IntVecIncubator {
	incu := IntVecIncubator{
		MaxIterCnt:      500,
		Population:      80,
		MutationRate:    0.05,
		IterCnt:         0,
		ChromLen:        10,
		TourRate:        0.9,
		TourSize:        5,
		MinChromEl:      0,
		MaxChromEl:      1,
		InitCallback:    func() {},
		EachCallback:    func() {},
		FinalCallback:   func() {},
		BestFitness:     math.Inf(-1),
		Best:            nil,
		SelectionMethod: "Replace",
		EncodingMethod:  "Binary",
		CrossoverMethod: "OnePoint",
		MutationMethod:  "FlipBits",
		ParentSelMethod: "Tournament",
	}
	return &incu
}

func (ivc *IntVecIncubator) Init() error {
	pool := make([]IntVecChromUber, ivc.Population)
	for i := 0; i < ivc.Population; i++ {
		//
		switch ivc.EncodingMethod {
		case "Binary":
			pool[i] = NewBitStrChromosome(ivc.ChromLen)
		case "Int":
			pool[i] = NewIntVecChromosome(ivc.ChromLen, ivc.MinChromEl, ivc.MaxChromEl)
		case "OrderList":
			pool[i] = NewOrderIntChromosome(ivc.MinChromEl, ivc.MaxChromEl)
		default:
			return errors.New("Encoding Method Not Found")
		}
	}
	ivc.Pool = pool
	return nil
}

func (ivc *IntVecIncubator) UpdateCount() {
	ivc.IterCnt += 1
}

func (ivc *IntVecIncubator) IsTerminal() bool {
	return ivc.IterCnt >= ivc.MaxIterCnt
}

func (ivc *IntVecIncubator) UpdateFitness() {
	for i, v := range ivc.Pool {
		ivc.Pool[i].SetFitness(ivc.FitnessFunc(v.Value()))
		if ivc.Pool[i].Fitness() > ivc.BestFitness {
			ivc.BestFitness = ivc.Pool[i].Fitness()
			ivc.Best = ivc.Pool[i].Copy().(IntVecChromUber)
		}
	}
}

func (ivc *IntVecIncubator) DoCrossover() error {
	n := ivc.Population / 2
	pool := make([]IntVecChromUber, 2*n)
	for i := 0; i < n; i++ {
		var pa IntVecChromUber
		var pa2 IntVecChromUber
		var ch IntVecChromUber
		var ch2 IntVecChromUber

		switch ivc.ParentSelMethod {
		case "Tournament":
			pa, pa2 = tourSelectPair(ivc.Pool, ivc.TourSize, ivc.TourRate)
			_ch, _ch2 := DoOnePointXover(pa, pa2)
			ch = _ch.(IntVecChromUber)
			ch2 = _ch2.(IntVecChromUber)
		default:
			return errors.New("No Matching ParentSelMethod Found in DoCrossover")
		}
		pool[2*i] = ch
		pool[2*i+1] = ch2
	}
	ivc.orldPool = ivc.Pool
	ivc.Pool = pool
	return nil
}

func (ivc *IntVecIncubator) DoSelection() error {
	switch ivc.SelectionMethod {
	case "Replace":
		//ivc.Pool = ivc.Offsprings
	default:
		return errors.New("No Matching SelectionMethod Found")
	}
	return nil
}

func (ivc *IntVecIncubator) DoMutation() error {
	p := ivc.MutationRate
	max := ivc.MaxChromEl
	min := ivc.MinChromEl
	for i, v := range ivc.Pool {
		switch ivc.MutationMethod {
		case "FlipBits":
			bv := v.(*BitStrChromosome)
			bv.Mutate(p)
			ivc.Pool[i] = bv
		case "RandomInt":
			iv := v.(*IntVecChromosome)
			iv.Mutate(p, min, max)
			ivc.Pool[i] = iv
		case "Swap":
		default:
			return errors.New("No Matching MutationMethod Found")
		}
	}
	return nil
}

func (ivc *IntVecIncubator) InitCb()  { ivc.InitCallback() }
func (ivc *IntVecIncubator) EachCb()  { ivc.EachCallback() }
func (ivc *IntVecIncubator) FinalCb() { ivc.FinalCallback() }
