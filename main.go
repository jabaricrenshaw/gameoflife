package main

import "github.com/jabaricrenshaw/gameoflife/types"

func main(){
    g := gameoflife.Game{
        Target_fps: 5,
    }

    g.Start()
} 

