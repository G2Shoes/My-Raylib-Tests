package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TILE   = 32 // Image size (32x32)
	WIDTH  = 512 // Window size (color map image size x 32, 16 x 32 = 512, where 32 is the size of the tile)
	HEIGHT = 512 // Window size (color map image size x 32, 16 x 32 = 512, where 32 is the size of the tile)
)

func Generate(grass *[]Floor, wall *[]Wall) {

	file, err := os.Open("sprites/map.png")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	img, _, err1 := image.Decode(file)

	if err1 != nil {
		fmt.Println(err1)
	}

	boundsInt := img.Bounds()

	for y := boundsInt.Min.Y; y < boundsInt.Max.Y; y++ {
		for x := boundsInt.Min.X; x < boundsInt.Max.X; x++ {

			c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)

			if c.R == 0 && c.G == 0 && c.B == 0 && c.A == 255 {
				newX := float32(x * TILE) // 1 * 32
				newY := float32(y * TILE) // 1 * 32

				*wall = append(*wall, Wall{
					position: rl.NewVector2(newX, newY),
				})
			}

			if c.R == 255 && c.G == 0 && c.B == 0 && c.A == 255 {
				newX := float32(x * TILE)// 1 * 32
				newY := float32(y * TILE)// 1 * 32

				*grass = append(*grass, Floor{
					position: rl.NewVector2(newX, newY),
				})
			}
		}
	}
}

func Paint(wal rl.Texture2D, gra rl.Texture2D, floor []Floor, walls []Wall) {
	for _, f := range floor {
		rl.DrawTextureV(gra, f.position, rl.RayWhite)
	}

	for _, w := range walls {
		rl.DrawTextureV(wal, w.position, rl.RayWhite)
	}
}

func main() {
	rl.InitWindow(WIDTH, HEIGHT, "RayGen")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	Init() // Debug

	var grass []Floor
	var walls []Wall

	wallTex := InitTexture("sprites/wall.png")
	floorTex := InitTexture("sprites/floor.png")

	Generate(&grass, &walls)

	defer DisposeTexture(wallTex, floorTex)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		Paint(wallTex, floorTex, grass, walls)
		DrawTextAnchored("Color Map Test", 0, 25, rl.RayWhite)
		rl.EndDrawing()
	}
}

func InitTexture(filepath string) rl.Texture2D {
	img := rl.LoadImage(filepath)
	tex := rl.LoadTextureFromImage(img)

	rl.UnloadImage(img)

	return tex
}

func DisposeTexture(tex1 rl.Texture2D, tex2 rl.Texture2D) {
	rl.UnloadTexture(tex1)
	rl.UnloadTexture(tex2)
}
