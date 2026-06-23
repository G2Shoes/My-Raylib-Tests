package main

import rl "github.com/gen2brain/raylib-go/raylib"

var BFrameCount int = 100

type Bird struct {
	Xpos     float32
	Ypos     float32
	Velocity float32
	Texture  rl.Texture2D
}

const (
	Gravity   = 0.315
	JumpForce = -5
)

func DrawTheBird(b Bird) {
	rl.DrawTexture(b.Texture, int32(b.Xpos), int32(b.Ypos), rl.RayWhite)
}

func BirdUpdate(b *Bird) {
	b.Velocity += Gravity

	if rl.IsKeyPressed(rl.KeySpace) {
		b.Velocity = JumpForce
	}

	b.Ypos += b.Velocity

	if b.Ypos > float32(rl.GetScreenHeight()+10) {
		gameover = true
	}

	if b.Ypos < 10 {
		gameover = true
	}
}
