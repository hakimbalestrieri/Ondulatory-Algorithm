package Ondulatoire

type Ondulatoire struct {
	moi int
	n	int
	voisins []bool
	nbVoisins int
	topologie [][]bool
}

func NewOndulatoire(n int) *Ondulatoire {
	o := new(Ondulatoire)
	o.n = n
	o.topologie := make([][]bool, n)
	for i := 0; i < n; i++ {
		o.topologie[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			o.topologie[i][j] = false
		}
	}
	for i := 0; i < n; i++ {
		o.topologie[o.moi][i] = o.voisins[i]
	}
	return &Ondulatoire{o.moi, n, o.voisins, o.nbVoisins, o.topologie}
}