package main

import ray "github.com/gen2brain/raylib-go/raylib"

func main() {
	ray.InitWindow(800, 400, "Krayzee Pong")
	defer ray.CloseWindow()
	ray.SetTargetFPS(30)

	sw := ray.GetScreenWidth()
	sh := ray.GetScreenHeight()

	p1 := Player{
		Xpos:   (sw / 2) - (100 / 2),
		Ypos:   sh - 35,
		Width:  100,
		Height: 10,
		Speed:  7,
	}

	b1 := Ball{
		Xpos:   400,
		Ypos:   200,
		Radius: 10,
		Color:  ray.Green,
		Speed:  7,
	}

	for !ray.WindowShouldClose() {
		ray.BeginDrawing()

		ray.ClearBackground(ray.Black)

		PaintThings(p1, b1)

		Update(&p1, &b1)

		ray.EndDrawing()
	}
}

func Update(p *Player, b *Ball) {
	KeyHandler(p)
	BallFunc(b)
}
