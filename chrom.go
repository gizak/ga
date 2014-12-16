package ga

import "math/rand"

type Chromosome interface {
	Value() interface{}
	Copy() Chromosome
	Age() int
	SetAge(int)
	Fitness() float64
	SetFitness(float64)
}

type emptyChrom struct {
	age     int
	fitness float64
}

func (c emptyChrom) Age() int {
	return c.age
}

func (c *emptyChrom) SetAge(a int) {
	c.age = a
}

func (c emptyChrom) Fitness() float64 {
	return c.fitness
}

func (c *emptyChrom) SetFitness(f float64) {
	c.fitness = f
}

type intChrom struct {
	emptyChrom
	value []int
}

func (ic intChrom) Value() interface{} {
	return ic.value
}

func (ic intChrom) Copy() Chromosome {
	nic := intChrom{}
	nic.age = ic.age
	nic.fitness = ic.fitness

	ns := make([]int, len(ic.value))
	copy(ns, ic.value)
	nic.value = ns

	return &nic
}

func newBinChrom(l int) Chromosome {
	ic := intChrom{}
	ic.value = make([]int, l)
	for i := range ic.value {
		if rand.Float32() > 0.5 {
			ic.value[i] = 1
		}
	}
	return &ic
}

// max inclusive
func newIntChrom(l, max int) Chromosome {
	ic := intChrom{}
	ic.value = make([]int, l)
	for i := range ic.value {
		ic.value[i] = rand.Intn(max + 1)
	}
	return &ic
}

func newOrderedIntChrom(l int) Chromosome {
	ic := intChrom{}
	ic.value = rand.Perm(l)
	return &ic
}
