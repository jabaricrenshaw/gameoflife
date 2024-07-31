package gameoflife

import (
	"fmt"
)

type Container struct {
	grid   [][]Cell
	width  int16
	height int16
}

func (ct Container) String() (s string) {
	for _, v := range ct.grid {
		s += fmt.Sprintf("%v\n", v)
	}
	return
}

func (ct *Container) update() {
	set_list := make([][]bool, ct.height)

	for i := range ct.height {
		set_list[i] = make([]bool, ct.width)
		for j := range ct.width {
			active_neighbors := 0

			for _, v := range ct.grid[i][j].neighbors {
				if v.active {
					active_neighbors += 1
				}
			}

			if active_neighbors == 3 {
				set_list[i][j] = true
			} else if ct.grid[i][j].active {
				if active_neighbors < 2 || active_neighbors > 3 {
					set_list[i][j] = false
				} else {
					set_list[i][j] = true
				}
			}
		}
	}

	for i := range set_list {
		for j := range set_list[0] {
			ct.grid[i][j] = Cell{
				active:    set_list[i][j],
				neighbors: ct.grid[i][j].neighbors,
			}
		}
	}
}

func NewContainer(width, height int16) *Container {
	return &Container{
		width:  width,
		height: height,
		grid:   make_grid(width, height),
	}
}
