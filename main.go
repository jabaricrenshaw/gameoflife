package main

import gameoflife "github.com/jabaricrenshaw/gameoflife/types"

func main() {
	// 5 FPS is reasonable for scrolling text
	g := gameoflife.Game{
		Target_fps: 5,
	}

	g.Start()
}
