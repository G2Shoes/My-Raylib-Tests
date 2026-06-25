package main

import rl "github.com/gen2brain/raylib-go/raylib"

const SPEED = 5

type Player struct {
	Color    rl.Color
	Position rl.Vector2
	Size     rl.Vector2
}

func Move(p *Player) {
	if rl.IsKeyDown(rl.KeyA) {
		p.Position.X -= SPEED
	}

	if rl.IsKeyDown(rl.KeyD) {
		p.Position.X += SPEED
	}
}
