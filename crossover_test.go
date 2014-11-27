package ga

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestOnePoint(t *testing.T) {
	Convey("Test OnePoint Crossover", t, func() {
		ch, ch2 := NewBitStrChromosome(10), NewBitStrChromosome(10)
		p := 5
		nch, nch2 := OnePointXover(ch, ch2, p)
		So(nch.Value()[:5], ShouldResemble, ch.Value()[:5])
		So(nch.Value()[5:], ShouldResemble, ch2.Value()[5:])
		So(len(nch.Value()), ShouldEqual, len(nch2.Value()))
		Convey("Can also be invoked by IntChrom", func() {
			ich, ich2 := NewIntVecChromosome(10, 0, 5), NewIntVecChromosome(10, 0, 5)
			nich, nich2 := OnePointXover(ich, ich2, 5)
			So(nich, ShouldHaveSameTypeAs, ich)
			So(nich2, ShouldHaveSameTypeAs, ich2)
		})
	})
}
