package main

import rl "github.com/gen2brain/raylib-go/raylib"

func GameOver() {
	rl.ClearBackground(rl.Black)
	DrawTextAnchored("GAME OVER", 2, 20, rl.Red)
}
