package ga

import "testing"
import "github.com/davecgh/go-spew/spew"

func TestIncuInit(t *testing.T) {
	cfg := NewDefaultConfg()
	cfg.ChromLen = 10
	inc := NewIncubator(cfg)
	inc.Init()
}

func TestIncuNext(t *testing.T) {
	cfg := NewDefaultConfg()
	cfg.ChromLen = 50
	cfg.MaxIterCnt = 100
	cfg.FitnessFunc = func(ch Chromosome) float64 {
		sl := ch.Value().([]int)
		sum := 0.0
		for _, v := range sl {
			sum += float64(v)
		}
		return sum
	}
	inc := NewIncubator(cfg)
	inc.Init()
	for inc.Next() {
		spew.Printf("%v,%v\n", inc.best.Fitness(), inc.AverageFitness())
	}
}
