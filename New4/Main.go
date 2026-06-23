package main

import (
	mrand "math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 400
const HEIGHT = 300

var gameended bool = false

func main() {
	rl.InitWindow(WIDTH, HEIGHT, "Falling Stuff Game")
	defer rl.CloseWindow()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(120)

	sw := rl.GetScreenWidth()
	sh := rl.GetScreenHeight()

	player := plr{
		Xpos:       (sw / 2) - (2.0 / 2),
		Ypos:       sh - 50,
		Speed:      3,
		SquareProp: 25,
		c:          rl.RayWhite,
	}

	ball := Projectile{
		Speed:  4,
		Radius: 7.25,
		c:      rl.RayWhite,
		posY:   int32(sh - 600),
		posX:   int32(mrand.Intn(sw)),
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Red)

		if !gameended {
			DrawTextAnchored("Falling Stuff Game Demo", 0, 15, rl.RayWhite)
			DrawTextAnchored("By G2S", 3, 15, rl.RayWhite)

			Draw(player)

			DrawProjectiles(ball)

			Update(&player, &ball)
		} else {
			GameOver()
		}

		rl.EndDrawing()
	}
}

func Update(p *plr, b *Projectile) {
	Move(p)

	ProjectileFall(b)

	if CheckCollision(*p, *b) {
		PlayOneShot("snd/error.wav", 1)
		gameended = true
	}
}

func CheckCollision(p plr, b Projectile) bool {
	ballCenter := rl.NewVector2(float32(b.posX), float32(b.posY))

	playerRect := rl.NewRectangle(
		float32(p.Xpos),
		float32(p.Ypos),
		float32(p.SquareProp),
		float32(p.SquareProp),
	)

	return rl.CheckCollisionCircleRec(ballCenter, b.Radius, playerRect)
}
