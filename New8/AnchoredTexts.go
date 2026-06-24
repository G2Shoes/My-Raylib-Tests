package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	anchorTopLeft     = 0
	anchorBottomLeft  = 1
	anchorCenter      = 2
	anchorTopRight    = 3
	anchorBottomRight = 4
)

func DrawTextAnchored(text string, anchor int, fontsize int32, color rl.Color) {
	var x int32
	var y int32

	textWidth := rl.MeasureText(text, fontsize)
	screenWidth := int32(rl.GetScreenWidth())
	screenHeight := int32(rl.GetScreenHeight())

	switch anchor {
	case anchorTopLeft:
		x = 10
		y = 10

	case anchorBottomLeft:
		x = 10
		y = screenHeight - fontsize - 10

	case anchorCenter:
		x = (screenWidth / 2) - (textWidth / 2)
		y = (screenHeight / 2) - (fontsize / 2)

	case anchorTopRight:
		x = screenWidth - textWidth - 10
		y = 10

	case anchorBottomRight:
		x = screenWidth - textWidth - 10
		y = screenHeight - fontsize - 10
	}
	rl.DrawText(text, x, y, fontsize, color)
}
