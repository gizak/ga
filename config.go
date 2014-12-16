package ga

type method uint

type encodingMethod struct {
	Binary   method
	Int      method
	OrderInt method
}

type selectionMethod struct {
	Truncate method
	Replace  method
}

type crossoverMethod struct {
	PMX      method
	onePoint method
	twoPoint method
	Order    method
	Position method
}

type mutationMethod struct {
	FlipBits  method
	RandomInt method
	Swap      method
	Shuffle   method
}

type parentSelMethod struct {
	Tournament method
	Rank       method
}

const (
	//encoding
	binary method = iota
	intVec
	orderedIntVec
	//parentSel
	tour
	rank
	//selection
	trunc
	replace
	//crossover
	pmx
	onePoint
	twoPoint
	order
	position
	//mutation
	flipBits
	randomInt
	swap
	shuffle
)

var Method = struct {
	EncodingMethod  encodingMethod
	ParentSelMethod parentSelMethod
	SelectionMethod selectionMethod
	MutationMethod  mutationMethod
	CrossoverMethod crossoverMethod
}{
	encodingMethod{binary, intVec, orderedIntVec},
	parentSelMethod{tour, rank},
	selectionMethod{trunc, replace},
	mutationMethod{flipBits, randomInt, swap, shuffle},
	crossoverMethod{pmx, onePoint, twoPoint, order, position}}

type Config struct {
	FitnessFunc     func(Chromosome) float64
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
	CrossoverRate   float64
	SelectionMethod method
	EncodingMethod  method
	CrossoverMethod method
	MutationMethod  method
	ParentSelMethod method
}

func NewDefaultConfg() Config {
	cfg := Config{}
	cfg.Population = 80
	cfg.MaxIterCnt = 500
	cfg.MutationRate = 0.03
	cfg.CrossoverRate = 0.9
	cfg.TourRate = 0.9
	cfg.TourSize = 5

	cfg.ParentSelMethod = Method.ParentSelMethod.Tournament
	cfg.MutationMethod = Method.MutationMethod.FlipBits
	cfg.SelectionMethod = Method.SelectionMethod.Replace
	cfg.CrossoverMethod = Method.CrossoverMethod.onePoint
	cfg.EncodingMethod = Method.EncodingMethod.Binary

	return cfg
}
