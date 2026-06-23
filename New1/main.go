package main

import rl "github.com/gen2brain/raylib-go/raylib"

const width = 800
const height = 600

func main() {
	rl.InitWindow(width, height, "Window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		Draw()
		Update()
	}
}

func Draw() {

	rl.BeginDrawing()
	rl.ClearBackground(rl.Blue)
	DrawTextAnchored("Hello World!", 2, 25, rl.RayWhite)
	rl.EndDrawing()
}

func Update() {
	CheckKeys()
}
