package main

import (
	ray "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Xpos   int
	Ypos   int
	Width  int32
	Height int32
	Speed  int32
}

type Ball struct {
	Xpos   int32
	Ypos   int32
	Radius float32
	Color  ray.Color
	Speed  int32
}

const (
	anchorTopLeft     = 0
	anchorBottomLeft  = 1
	anchorCenter      = 2
	anchorTopRight    = 3
	anchorBottomRight = 4
)

func PaintThings(p Player, b Ball) {
	ray.DrawCircle(b.Xpos, b.Ypos, b.Radius, b.Color)
	DrawTextAnchored("KRAYZEE PONG", 0, 25, ray.Green)
	DrawTextAnchored("0", 3, 25, ray.Green)
	ray.DrawRectangle(int32(p.Xpos), int32(p.Ypos), p.Width, p.Height, ray.Green)
}

func KeyHandler(p *Player) {
	if ray.IsKeyDown(ray.KeyA) {
		p.Xpos -= int(p.Speed)
	}
	if ray.IsKeyDown(ray.KeyD) {
		p.Xpos += int(p.Speed)
	}
}

func DrawTextAnchored(text string, anchor int, fontsize int32, color ray.Color) {
	var x int32
	var y int32

	tw := ray.MeasureText(text, fontsize)

	sw := int32(ray.GetScreenWidth())
	sh := int32(ray.GetScreenHeight())

	switch anchor {
	case anchorTopLeft:
		x = 10
		y = 10
	case anchorBottomLeft:
		x = 10
		y = sh - fontsize - 10
	case anchorCenter:
		x = (sw / 2) - (tw / 2)
		y = (sh / 2) - (fontsize / 2)
	case anchorTopRight:
		x = sw - 15
		y = 10
	case anchorBottomRight:
		x = sw - 10
		y = sh - 10
	}

	ray.DrawText(text, x, y, fontsize, color)
}
