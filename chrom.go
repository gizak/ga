package ga

import "math/rand"

// Chromosome Representation
type IntVecChromosome struct {
	age     int
	value   []int
	fitness float64
}

//BitStrChromosome is 0-1 represented
type BitStrChromosome struct{ IntVecChromosome }

// OrderIntChromosome for TSP
type OrderIntChromosome struct{ IntVecChromosome }

func (vc *IntVecChromosome) Copy() IntVecChromValuer {
	nvc := *vc
	return &nvc
}

func (vc IntVecChromosome) Value() []int {
	return vc.value
}

func (vc IntVecChromosome) Age() int {
	return vc.age
}

func (vc IntVecChromosome) Fitness() float64 {
	return vc.fitness
}

func (vc *IntVecChromosome) SetValue(sl []int) {
	vc.value = sl
}

func (vc *IntVecChromosome) SetFitness(f float64) {
	vc.fitness = f
}

func (vc *IntVecChromosome) SetAge(n int) {
	vc.age = n
}

type IntVecChromValuer interface {
	Value() []int
	//Fitness() float64
	//Age() int
	SetValue([]int)
	//SetFitness(float64)
	//SetAge(int)
	Copy() IntVecChromValuer
}

type IntVecChromUber interface {
	IntVecChromValuer
	ChromFitnessTester
	ChromAger
}

type ChromFitnessTester interface {
	Fitness() float64
	SetFitness(float64)
}

type ChromAger interface {
	Age() int
	SetAge(int)
}

//
func NewBitStrChromosome(n int) *BitStrChromosome {
	intSl := make([]int, n)
	for i := 0; i < n; i++ {
		if rand.Float32() >= 0.5 {
			intSl[i] = 1
		} else {
			intSl[i] = 0
		}
	}

	chrom := BitStrChromosome{}
	chrom.value = intSl
	return &chrom
}

func (bc *BitStrChromosome) Copy() IntVecChromValuer {
	nbc := new(BitStrChromosome)
	nbc.IntVecChromosome = *bc.IntVecChromosome.Copy().(*IntVecChromosome)
	return nbc
}

//
func NewIntVecChromosome(n, min, max int) *IntVecChromosome {
	intSl := make([]int, n)
	for i := 0; i < n; i++ {
		intSl[i] = rand.Intn(max - min + 1)
		intSl[i] += min
	}
	chrom := IntVecChromosome{}
	chrom.value = intSl
	return &chrom
}

//
func NewOrderIntChromosome(min, max int) *OrderIntChromosome {
	intSl := rand.Perm(max - min + 1)
	for i := range intSl {
		intSl[i] += min
	}
	chrom := OrderIntChromosome{}
	chrom.value = intSl
	return &chrom
}
