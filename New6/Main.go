package main

import (
	mrand "math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 800
const HEIGHT = 640

var gameover bool = false

func main() {
	rl.InitWindow(WIDTH, HEIGHT, "survival game tech demo")
	defer rl.CloseWindow()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)

	sw := rl.GetScreenWidth()
	sh := rl.GetScreenHeight()

	preferred_plrsize := 45

	img := rl.LoadImage("sprites/enemy.png")
	tex := rl.LoadTextureFromImage(img)
	defer rl.UnloadImage(img)
	defer rl.UnloadTexture(tex)

	plr := Player{
		Xpos:        (int32(sw/2) - (int32(preferred_plrsize) / 2)),
		Ypos:        (int32(sh - 75)),
		Speed:       5,
		ColorTop:    rl.Red,
		ColorBottom: rl.SkyBlue,
		Size:        float32(preferred_plrsize),
	}

	var enemy []Enemy
	var frameCount int32

	var pellet []Pellet

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if !gameover {
			rl.ClearBackground(rl.Black)
			DrawPlayer(plr)
			rectP := rl.NewRectangle(float32(plr.Xpos), float32(plr.Ypos), float32(preferred_plrsize), float32(preferred_plrsize))

			for _, e := range enemy {
				DrawEnemy(e, tex)
			}

			for _, p := range pellet {
				DrawPellet(p)
			}

			Update(&plr, plr, &pellet, &enemy, &frameCount, rectP)
			DrawTextAnchored("Survival Game Tech Demo", 0, 25, rl.RayWhite)
		}

		if gameover {
			rl.ClearBackground(rl.Red)
			DrawTextAnchored("GAME OVER", 2, 50, rl.RayWhite)
		}

		rl.EndDrawing()
	}
}

func Update(p *Player, plrCopy Player, pel *[]Pellet, e *[]Enemy, frameCount *int32, plrRect rl.Rectangle) {
	Move(p)
	Fire(plrCopy, pel)

	for i := 0; i < len(*pel); i++ {
		FlyUp(&(*pel)[i])
		peRect := rl.NewRectangle(float32((*pel)[i].Xpos), float32((*pel)[i].Ypos), float32((*pel)[i].Radius*2), float32((*pel)[i].Radius*2))

		for j := 0; j < len(*e); j++ {
			enemyRect := rl.NewRectangle(float32((*e)[j].Xpos), float32((*e)[j].Ypos), 64, 64)

			if CheckEnemyPelletCollision(enemyRect, peRect) {
				*e = append((*e)[:j], (*e)[j+1:]...)
				*pel = append((*pel)[:i], (*pel)[i+1:]...)
				i--
				PlayOneShot("snd/boom.wav", 1)
				break
			}
		}
	}

	(*frameCount)++

	if *frameCount >= 60 {
		enemy := Enemy{
			Xpos:  int32(mrand.Intn(rl.GetScreenWidth()) + 1),
			Ypos:  -80,
			Speed: int32(11 / 2),
			Scale: 0.05,
		}

		*e = append(*e, enemy)
		*frameCount = 0
	}

	for j := 0; j < len(*e); j++ {
		Attack(&(*e)[j])
		enemyRect := rl.NewRectangle(float32((*e)[j].Xpos), float32((*e)[j].Ypos), 45, 45)

		if CheckEnemyPlayerCollision(plrRect, enemyRect) {
			gameover = true
			PlayOneShot("snd/error.wav", 1)
		}

		if (*e)[j].Ypos > int32(rl.GetScreenHeight()) {
			gameover = true
			PlayOneShot("snd/error.wav", 1)
		}
	}
}

func CheckEnemyPlayerCollision(pRect rl.Rectangle, eRect rl.Rectangle) bool {
	return rl.CheckCollisionRecs(pRect, eRect)
}

func CheckEnemyPelletCollision(eRect rl.Rectangle, peRect rl.Rectangle) bool {
	return rl.CheckCollisionRecs(eRect, peRect)
}
