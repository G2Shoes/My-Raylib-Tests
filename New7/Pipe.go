package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Pipes struct {
	BottomYpos int32
	TopYPos    int32
	TopXPos    int32
	Sprite     rl.Texture2D
	Active     bool
}

var speed float32 = 1.7

func DrawThePipes(pipes []Pipes) {
	for _, pipe := range pipes {
		if pipe.Active {
			rl.DrawTexture(pipe.Sprite, pipe.TopXPos, pipe.BottomYpos, rl.RayWhite)

			FlippedX := float32(pipe.TopXPos) + float32(pipe.Sprite.Width)

			rl.DrawTextureEx(pipe.Sprite, rl.NewVector2(FlippedX, float32(pipe.TopYPos)), 180, 1.0, rl.RayWhite)
		}
	}
}

func PipeMove(pipes []Pipes) {
	for i := range pipes {
		pipes[i].TopXPos -= int32(speed)

		if pipes[i].TopXPos < -50 {
			pipes[i].Active = false
		}
	}
}
