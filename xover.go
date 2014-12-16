package ga

import "math/rand"

func cpIntSl(a []int) []int {
	l := len(a)
	b := make([]int, l)
	copy(b, a)
	return b
}

func onePointXoverAt(c1, c2 []int, p int) ([]int, []int) {
	nc1 := make([]int, 0, len(c1))
	nc2 := make([]int, 0, len(c2))

	subc1A := c1[:p]
	subc1B := c1[p:]

	subc2A := c2[:p]
	subc2B := c2[p:]

	nc1 = append(nc1, subc1A...)
	nc1 = append(nc1, subc2B...)
	nc2 = append(nc2, subc2A...)
	nc2 = append(nc2, subc1B...)

	return nc1, nc2
}

func onePointXover(c1, c2 []int) ([]int, []int) {
	p := rand.Intn(len(c1))
	return onePointXoverAt(c1, c2, p)
}
