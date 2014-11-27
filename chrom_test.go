package ga

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestNewChrom(t *testing.T) {
	Convey("Structing New BitStr", t, func() {
		s := NewBitStrChromosome(10)

		So(len(s.Value()), ShouldEqual, 10)
		So(0, ShouldBeIn, s.Value())
		So(1, ShouldBeIn, s.Value())
		So(s.Value(), ShouldNotContain, 2)
		So(s.Age(), ShouldEqual, 0)
		So(s.Fitness(), ShouldEqual, 0)
	})
	Convey("Structing New IntChromosome", t, func() {
		s := NewIntVecChromosome(100, 1, 10)
		So(len(s.Value()), ShouldEqual, 100)
		So(1, ShouldBeIn, s.value)
		So(10, ShouldBeIn, s.Copy().Value())
		So(s.Value(), ShouldNotContain, 0)
		So(s.Age(), ShouldEqual, 0)
		So(s.Fitness(), ShouldEqual, 0)
	})
	Convey("Structing New OrderIntChromosome", t, func() {
		s := NewOrderIntChromosome(1, 100)
		So(len(s.Value()), ShouldEqual, 100)
		So(1, ShouldBeIn, s.Value())
		So(100, ShouldBeIn, s.Value())
		So(s.Value(), ShouldNotContain, 0)
		So(s.Age(), ShouldEqual, 0)
		So(s.Fitness(), ShouldEqual, 0)
		//ns := s.Copy()
		s.SetAge(100)
		So(s.Age(), ShouldEqual, 100)
		So(s.Age(), ShouldEqual, 100)
		//So(ns.Age(), ShouldEqual, 0)
	})
}
