package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Pellet struct {
	Speed  int32
	Xpos   int32
	Ypos   int32
	Radius int32
	Color  rl.Color
}

func DrawPellet(p Pellet) {
	rl.DrawCircle(p.Xpos, p.Ypos, float32(p.Radius), p.Color)
}

func FlyUp(p *Pellet) {
	p.Ypos -= p.Speed
}
