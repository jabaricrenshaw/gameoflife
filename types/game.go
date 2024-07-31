package gameoflife

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Container type will be modified every frame
// Options are set once on game init and only contain game "window size"
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
	sc := bufio.NewScanner(os.Stdin)

	fmt.Println(gme.container)
	for {
		var x_in, y_in int16 = -1, -1
		fmt.Println("Cell activation.")
		fmt.Printf("Provide X Y coordinates. (Within x=[0,%d] y=[0,%d]): ", gme.container.width-1, gme.container.height-1)

		sc.Scan()
		n, err := fmt.Sscanf(sc.Text(), "%d %d\n", &x_in, &y_in)
		if err == nil {
			if (x_in >= 0 && x_in < gme.container.width) && (y_in >= 0 && y_in < gme.container.height) {
				gme.container.grid[gme.container.height-y_in-1][x_in].active = !gme.container.grid[gme.container.height-y_in-1][x_in].active
				fmt.Printf("%s\n", gme.container)
			} else {
				fmt.Println("Invalid coordinates!\n")
			}
		} else if n == 0 && err.Error() == "EOF" {
			break
		} else {
			fmt.Printf("%s\n\n", getInputErr(err))
		}
	}
}

func get_opts() Options {
	sc := bufio.NewScanner(os.Stdin)
	var temp_w, temp_h int16

	for {
		temp_w, temp_h = -1, -1
		fmt.Printf("Term size? (WxH): ")

		sc.Scan()
		_, err := fmt.Sscanf(sc.Text(), "%d %d\n", &temp_w, &temp_h)
		if err != nil {
			fmt.Printf("%s\n\n", getInputErr(err))
		} else if temp_w <= 0 || temp_h <= 0 {
			fmt.Println("Invalid size!\n")
		} else {
			fmt.Println("")
			break
		}
	}

	return Options{
		width:  temp_w,
		height: temp_h,
	}
}

func getInputErr(err error) string {
	switch err.Error() {
	case "EOF":
		return "Not enough values!"
	case "newline in format does not match input":
		return "Too many values!"
	default:
		return "Bad input."
	}
}

func (gme *Game) Start() {
	windowOpts := get_opts()
	gme = &Game{
		container:     NewContainer(windowOpts.width, windowOpts.height),
		Target_fps:    gme.Target_fps,
		frame_counter: 0,
	}

	gme.init_cells()

	wait_time_ms := time.Second / time.Duration(gme.Target_fps)
	var total_time time.Duration
	for {
		// Incrementing frame counter and showing stats
		gme.frame_counter++
		fmt.Printf("Frames elapsed: %d (%.2f FPS @ %s/frame)\n%s", gme.frame_counter, gme.Target_fps, wait_time_ms, gme.container)

		// Timing this frame update
		start := time.Now()
		gme.container.update()
		fmt.Printf("(frame gen took: %s)\n\n", time.Since(start))

		// Waiting for next frame

		time.Sleep(wait_time_ms)
		total_time += (time.Since(start))
	}
}
