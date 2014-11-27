package ga

import "math/rand"

func OnePointXover(ch, ch2 IntVecChromValuer, p int) (IntVecChromValuer, IntVecChromValuer) {
	if p < 0 || p >= len(ch.Value()) {
		panic(ch.Value())
	}

	nch := ch.Copy()
	nch2 := ch2.Copy()

	subChA := ch.Value()[:p]
	subChB := ch.Value()[p:]
	subCh2A := ch2.Value()[:p]
	subCh2B := ch2.Value()[p:]

	nch.SetValue(make([]int, 0, len(ch.Value())))
	nch.SetValue(append(nch.Value(), subChA...))
	nch.SetValue(append(nch.Value(), subCh2B...))

	nch2.SetValue(make([]int, 0, len(ch2.Value())))
	nch2.SetValue(append(nch2.Value(), subCh2A...))
	nch2.SetValue(append(nch2.Value(), subChB...))

	return nch, nch2
}

func DoOnePointXover(ch, ch2 IntVecChromValuer) (IntVecChromValuer, IntVecChromValuer) {
	l := len(ch.Value())
	p := rand.Intn(l)
	return OnePointXover(ch, ch2, p)
}
