package main

import (
	mrand "math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Projectile struct {
	Speed  int32
	Radius float32
	c      rl.Color
	posY   int32
	posX   int32
}

func DrawProjectiles(pr Projectile) {
	rl.DrawCircle(pr.posX, pr.posY, float32(pr.Radius), pr.c)
}

func ProjectileFall(pr *Projectile) {
	pr.posY += pr.Speed

	if pr.posY > int32(rl.GetScreenHeight()) {
		pr.posY = 0
		pr.posX = int32(mrand.Intn(rl.GetScreenWidth()))
		PlayOneShot("snd/click.wav", 1)
	}
}
