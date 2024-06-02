package main

import (
    "gameoflife/types"
)

// Structures
/************************************************************/
/*
type Cell struct {
    active	bool
    neighbors	[]*Cell
}

type Grid [][]Cell
*/

/*
type Container struct {
    grid    Grid
    width   int
    height  int
}
*/

/*
type Game struct {
    container	    *Container
    opts            *Options
    target_fps	    float64
    frame_counter   int
}
*/
/*
type Options struct {
    width   int
    height  int
}
*/
/************************************************************/



// Cell receivers
/************************************************************/
/*
func (c Cell) String() string {
    if c.active {
	return "*"
    }
    return "-"
}
*/
/************************************************************/



// Continer receivers and related
/************************************************************/
/*
func (cont Container) String() (s string) {
    for _, v := range cont.grid {
        s += fmt.Sprintf("%v\n", v);
    }
    return
}
*/

/*
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
*/

/*
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
*/

/*
func NewContainer(width, height int) (cont *Container) {
    return &Container {
        width: width,
        height: height,
        grid: make_grid(width, height),
    }
}
*/
/************************************************************/
/*
func (gme *Game) init_cells(){
    sc := bufio.NewScanner(os.Stdin)
    for {
        x_in, y_in := -1, -1
        fmt.Printf("Cell activation.\n")
        fmt.Printf("Provide X Y coordinates. (Within %d %d): ", gme.container.width, gme.container.height)
        
        sc.Scan()
        
        if n, err := fmt.Sscanf(sc.Text(), "%d %d\n", &x_in, &y_in); n == 0 {
            break
        } else if err == nil && ((x_in >= 0 && x_in < gme.container.width) && (y_in >= 0 && y_in < gme.container.height)) {
            gme.container.grid[gme.container.height - y_in - 1][x_in].active = !gme.container.grid[gme.container.height - y_in - 1][x_in].active
            fmt.Printf("%s\n", gme.container)
        } else if err != nil && err.Error() == "EOF" {
            fmt.Printf("Not enough values!\n\n")
        } else if err != nil && err.Error() == "newline in format does not match input" {
            fmt.Printf("Too many values!\n\n")
        } else if err != nil {
            fmt.Printf("Bad input.\n\n")
        } else {
            fmt.Printf("Invalid coordinates! Out of bounds.\n\n")
        }
    }
}
*/

/*
func (gme *Game) get_opts() {
    sc := bufio.NewScanner(os.Stdin)
    
    for {
        fmt.Printf("Term size? (WxH): ")
        sc.Scan() 
        temp_w, temp_h := -1, -1
        if _, err := fmt.Sscanf(sc.Text(), "%d %d\n", &temp_w, &temp_h); err != nil {
            fmt.Printf("Bad input.\n\n")
        } else {
            gme.opts = &Options {
                width: temp_w,
                height: temp_h,
            }
            break
        }
    }
}
*/

/*
func (gme *Game) Start() {
    gme.get_opts()
    gme.container = NewContainer(gme.opts.width, gme.opts.height)
    gme.init_cells()
    
    wait_time_ms := (1000.00/gme.target_fps)
    for {
        gme.frame_counter++
        fmt.Printf("Frames elapsed: %d\n%.2f FPS (%.2f ms/frame)\n%s\n", gme.frame_counter, gme.target_fps, wait_time_ms, gme.container)
        gme.container.update()
        time.Sleep(time.Duration(wait_time_ms) * time.Millisecond)
    }
}
*/

func main(){
    g := Game{
        target_fps: 5,
    }

    g.Start()
} 

