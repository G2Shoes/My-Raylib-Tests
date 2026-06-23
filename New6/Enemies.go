package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Enemy struct {
	Xpos  int32
	Ypos  int32
	Scale float32
	Speed int32
}

func DrawEnemy(e Enemy, sprite rl.Texture2D) {
	rl.DrawTexture(sprite, e.Xpos, e.Ypos, rl.RayWhite)
}

func Attack(e *Enemy) {
	e.Ypos += e.Speed
}
