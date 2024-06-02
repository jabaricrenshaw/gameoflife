package gameoflife

import (
    "fmt"
    "bufio"
    "os"
    "time"
)

type Game struct {
    container	    *Container
    opts            *Options
    Target_fps	    float64
    frame_counter   int
}


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


func (gme *Game) Start() {
    gme.get_opts()
    gme.container = NewContainer(gme.opts.width, gme.opts.height)
    gme.init_cells()
    
    wait_time_ms := (1000.00/gme.Target_fps)
    for {
        gme.frame_counter++
        fmt.Printf("Frames elapsed: %d\n%.2f FPS (%.2f ms/frame)\n%s\n", gme.frame_counter, gme.Target_fps, wait_time_ms, gme.container)
        gme.container.update()
        time.Sleep(time.Duration(wait_time_ms) * time.Millisecond)
    }
}
