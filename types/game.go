package gameoflife

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// container     *Container 	will be modified every frame
// Target_fps    float64 		should be exported so that it can be adjusted
type Game struct {
	container     *Container
	frame_counter int
	Target_fps    float64
}

// int16 is a suitable type for our use case
type Options struct {
	width  int16
	height int16
}

func (gme *Game) init_cells() {
	fmt.Println(gme.container)
	for {
		var x_in, y_in int16 = -1, -1
		fmt.Println("Cell activation.")
		fmt.Printf("Provide X Y coordinates. (Within x=[0,%d] y=[0,%d]): ", gme.container.width-1, gme.container.height-1)

		n, err := fmt.Scanf("%d %d\n", &x_in, &y_in)
		if err == nil {
			if (x_in >= 0 && x_in < gme.container.width) && (y_in >= 0 && y_in < gme.container.height) {
				gme.container.grid[gme.container.height-y_in-1][x_in].active = !gme.container.grid[gme.container.height-y_in-1][x_in].active
				fmt.Printf("%s\n", gme.container)
			} else {
				fmt.Println("Invalid coordinates!\n")
			}
		} else if n == 0 && err.Error() == "unexpected newline" {
			break
		} else {
			fmt.Printf("%s\n\n", getInputErr(err))
		}
	}
}

func get_opts() Options {
	for {
		var temp_w, temp_h int16 = -1, -1
		fmt.Printf("Grid size? (WxH): ")

		n, err := fmt.Scanf("%d %d\n", &temp_w, &temp_h)
		if err != nil {
			fmt.Printf("%s\n\n", getInputErr(err))
		} else if temp_w <= 0 || temp_h <= 0 {
			fmt.Println("Invalid size!\n")
		} else if n == 2 {
			return Options{
				width:  temp_w,
				height: temp_h,
			}
		}
	}
}

func getInputErr(err error) string {
	switch err.Error() {
	case "unexpected newline":
		return "Not enough values!"
	case "newline in format does not match input":
		return "Too many values!"
	}
	return "Bad input."
}

func (gme *Game) Start() {
	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.Exit(0)
	}()

	windowOpts := get_opts()
	gme = &Game{
		container:     NewContainer(windowOpts.width, windowOpts.height),
		Target_fps:    gme.Target_fps,
		frame_counter: 0,
	}

	// Setting up game loop timer
	frame_wait_time := time.Second / time.Duration(gme.Target_fps)
	t := time.NewTicker(frame_wait_time)
	defer t.Stop()

	// Getting user input to activate cells
	gme.init_cells()

	// Main game loop
	for range t.C {
		// Incrementing counter and showing stats
		gme.frame_counter++
		fmt.Printf("Frames elapsed: %d (%.2f FPS @ %s/frame)\n%s", gme.frame_counter, gme.Target_fps, frame_wait_time, gme.container)

		// Timing this frame update
		start_frame_gen := time.Now()
		gme.container.update()
		fmt.Printf("(frame gen took: %s)\n\n", time.Since(start_frame_gen))
	}
}
