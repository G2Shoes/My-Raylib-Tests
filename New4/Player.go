package main

import rl "github.com/gen2brain/raylib-go/raylib"

type plr struct {
	Speed      int32
	Xpos       int
	Ypos       int
	SquareProp int32
	c          rl.Color
}

func Draw(p plr) {
	rl.DrawRectangle(int32(p.Xpos), int32(p.Ypos), p.SquareProp, p.SquareProp, p.c)
}

func Move(p *plr) {
	if rl.IsKeyDown(rl.KeyA) {
		p.Xpos -= int(p.Speed)
	}

	if rl.IsKeyDown(rl.KeyD) {
		p.Xpos += int(p.Speed)
	}
}
