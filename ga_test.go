package ga

import "testing"
import . "github.com/smartystreets/goconvey/convey"
import "github.com/davecgh/go-spew/spew"
import "encoding/csv"
import "os"
import "fmt"

func TestGAIncubator(t *testing.T) {
	ivc := NewIntVecIncubator()
	Convey("TestInit", t, func() {
		ivc.Population = 100
		ivc.MaxIterCnt = 100
		ivc.ChromLen = 50
		ivc.EncodingMethod = "Binary"
		ivc.FitnessFunc = func(sl []int) float64 {
			sum := 0.0
			for _, v := range sl {
				sum += float64(v)
			}
			return sum
		}
		ivc.Init()
		So(len(ivc.Pool), ShouldEqual, 100)
	})
	Convey("TestUpdate", t, func() {
		ivc.UpdateFitness()
		f, _ := os.Create("dump-init.text")
		defer f.Close()
		spew.Fprintf(f, "%#v", *ivc)
	})

	Convey("TestRunning", t, func() {
		f, _ := os.Create("test-running-fitness.csv")
		defer f.Close()
		rec := csv.NewWriter(f)

		ivc.InitCallback = func() { rec.Write([]string{"Cnt", "Fitness"}) }
		ivc.EachCallback = func() { rec.Write([]string{fmt.Sprint(ivc.IterCnt), fmt.Sprint(ivc.BestFitness)}) }
		ivc.FinalCallback = func() { rec.Flush() }

		EvolveOnce(ivc)

	})
}
