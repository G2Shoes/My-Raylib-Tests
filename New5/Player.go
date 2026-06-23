package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Speed   int32
	Xpos    int32
	Ypos    int32
	Texture rl.Texture2D
}

func DrawPlayer(p Player) {
	position := rl.NewVector2(float32(p.Xpos), float32(p.Ypos))

	rl.DrawTextureEx(p.Texture, position, 0, 1.0, rl.RayWhite)
}

func move(p *Player) {
	if rl.IsKeyDown(rl.KeyA) {
		p.Xpos -= p.Speed
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.Xpos += p.Speed
	}
	if rl.IsKeyDown(rl.KeyW) {
		p.Ypos -= p.Speed
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.Ypos += p.Speed
	}
}
