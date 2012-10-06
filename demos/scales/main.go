package main

import (
	"github.com/ajhager/rog"
	_ "github.com/ajhager/rog/glfw"
)

var drainbow = rog.Discrete(rog.Red, rog.Orange, rog.Yellow, rog.Green, rog.Blue, rog.Purple, rog.Magenta)
var lrainbow = rog.Linear(rog.Red, rog.Orange, rog.Yellow, rog.Green, rog.Blue, rog.Purple, rog.Magenta)

func main() {
	rog.Open(20, 10, 2, "rog", "../../data/font.png")
    rog.Set(0, 3, rog.Burn(rog.Grey), drainbow, "   Discrete Scale   ")
    rog.Set(0, 6, rog.Burn(rog.Grey), lrainbow, "    Linear Scale    ")
	for rog.IsOpen() {
		rog.Set(0, 9, nil, nil, "%v", rog.Fps())
		if rog.Key() == rog.Escape {
			rog.Close()
		}
		rog.Flush()
	}
}
