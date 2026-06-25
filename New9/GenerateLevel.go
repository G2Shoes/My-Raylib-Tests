package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const TILE = 32
const PLAYERSIZEX = 32
const PLAYERSIZEY = 32

func InitLevelGenerator(floor *[]Floor, sky *[]Sky, testBlocks *[]TestBlock, p *Player) {
	file, err := os.Open("Map/level.png")

	if err != nil {
		fmt.Println(err)
	}

	img, _, err1 := image.Decode(file)

	if err1 != nil {
		fmt.Println(err1)
	}

	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)

			if c.R == 255 && c.G == 0 && c.B == 0 && c.A == 255 {
				p.Position = rl.NewVector2(float32(x*PLAYERSIZEX), float32(y*PLAYERSIZEY))
				p.Size = rl.NewVector2(PLAYERSIZEX, PLAYERSIZEY)
				p.Color = rl.RayWhite
				PlayerSpawned = true

				newX := float32(x * TILE)
				newY := float32(y * TILE)

				*sky = append(*sky, Sky{
					Position: rl.NewVector2(newX, newY),
				})
			}

			if c.R == 0 && c.G == 0 && c.B == 0 && c.A == 255 {
				newX := float32(x * TILE)
				newY := float32(y * TILE)

				*sky = append(*sky, Sky{
					Position: rl.NewVector2(newX, newY),
				})
			}

			if c.R == 255 && c.G == 255 && c.B == 255 && c.A == 255 {
				floorX := float32(x * TILE)
				floorY := float32(y * TILE)

				*floor = append(*floor, Floor{
					Position: rl.NewVector2(floorX, floorY),
				})
			}

			if c.R == 100 && c.G == 255 && c.B == 0 && c.A == 255 {
				newX1 := float32(x * TILE)
				newY1 := float32(y * TILE)

				*testBlocks = append(*testBlocks, TestBlock{
					Position: rl.NewVector2(newX1, newY1),
				})
			}
		}
	}
}

func PaintLevel(skySlice []Sky, floorSlice []Floor, floorTex rl.Texture2D, skyTex rl.Texture2D, tblock []TestBlock, tBlockTex rl.Texture2D) {
	for _, s := range skySlice {
		rl.DrawTextureV(skyTex, s.Position, rl.White)
	}

	for _, f := range floorSlice {
		rl.DrawTextureV(floorTex, f.Position, rl.White)
	}

	for _, tb := range tblock {
		rl.DrawTextureV(tBlockTex, tb.Position, rl.White)
	}
}
