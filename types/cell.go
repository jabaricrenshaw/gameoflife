package gameoflife

type Grid [][]Cell

type Cell struct {
    active	bool
    neighbors	[]*Cell
}


func (c Cell) String() string {
    if c.active {
	return "*"
    }
    return "-"
}


func make_grid(width, height int) Grid {
    g := make([][]Cell, height)
    for i := range height {
        g[i] = make([]Cell, width)
    }

    for i := range height {
        for j := range width {
            for a := -1; a <= 1; a++ {
                for b := -1; b <= 1; b++ {
                    if !(a == 0 && b == 0) && 
                        (a + i >= 0 && a + i < height) &&
                        (b + j >= 0 && b + j < width) {	
                        g[i][j].neighbors = append(g[i][j].neighbors, &g[a + i][b + j])
                    }
                }
            }
        }
    }
    
    return g
}
