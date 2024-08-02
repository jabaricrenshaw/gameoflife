package gameoflife

type Cell struct {
	active    bool
	neighbors []*Cell
}

func (ce Cell) String() string {
	if ce.active {
		return "*"
	}
	return "-"
}

func make_grid(width, height int16) [][]Cell {
	g := make([][]Cell, height)
	for i := range g {
		g[i] = make([]Cell, width)
	}

	var vert_neighbor, hor_neighbor int16
	for i := range height {
		for j := range width {
			for vert_neighbor = -1; vert_neighbor <= 1; vert_neighbor++ {
				for hor_neighbor = -1; hor_neighbor <= 1; hor_neighbor++ {
					if !(vert_neighbor == 0 && hor_neighbor == 0) {
						g[i][j].neighbors = append(g[i][j].neighbors, &g[((vert_neighbor+i)%height+height)%height][((hor_neighbor+j)%width+width)%width])
					}
				}
			}
		}
	}

	return g
}
