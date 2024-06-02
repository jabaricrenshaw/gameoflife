package gameoflife

import (
    "fmt"
)

type Container struct {
    grid    Grid
    width   int
    height  int
}


func (cont Container) String() (s string) {
    for _, v := range cont.grid {
        s += fmt.Sprintf("%v\n", v);
    }
    return
}


func (cont *Container) update() {
    new_grid := make_grid(cont.width, cont.height)
   
    for i := range cont.height {
        for j := range cont.width {
            active_count := 0
            for _, v := range (*cont).grid[i][j].neighbors {
                if v.active {
                    active_count++
                }
            }
           
            if (*cont).grid[i][j].active {
                if active_count <= 1 || active_count >= 4 {
                    new_grid[i][j].active = false
                } else {
                    new_grid[i][j].active = true
                }
            } else if active_count == 3 {
                new_grid[i][j].active = true
            }

        }
    }

    cont.grid = new_grid
}


func NewContainer(width, height int) (cont *Container) {
    return &Container {
        width: width,
        height: height,
        grid: make_grid(width, height),
    }
}
