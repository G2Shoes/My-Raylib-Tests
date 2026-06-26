package main

import (
	"encoding/json"
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	Width    int32     `json:"Width"`
	Height   int32     `json:"Height"`
	TileSize int32     `json:"TileSize"`
	Data     [][]int32 `json:"Data"`
}

func LoadLevel(filepath string) Level {
	file, err := os.ReadFile("level.json")
	if err != nil {
		fmt.Println(err)
	}

	var level Level

	err1 := json.Unmarshal(file, &level)

	if err1 != nil {
		fmt.Println(err1)
	}

	return level
}

func main() {

	rl.InitWindow(1080, 800, "JSON parse test")
	defer rl.CloseWindow()

	levelToLoad := LoadLevel("level.json")

	offsetX := (1080 - (levelToLoad.Width * levelToLoad.TileSize)) / 2
	offsetY := (800 - (levelToLoad.Height * levelToLoad.TileSize)) / 2

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)

		for y, rowH := range levelToLoad.Data {
			for x, tile := range rowH {
				posX := int32(x)*levelToLoad.TileSize + offsetX
				posY := int32(y)*levelToLoad.TileSize + offsetY

				if tile == 1 {
					rl.DrawRectangleV(rl.NewVector2(float32(posX), float32(posY)), rl.NewVector2(float32(levelToLoad.TileSize), float32(levelToLoad.TileSize)), rl.RayWhite)
				}
			}
		}

		DrawTextAnchored("JSON Level Loader test", 0, 25, rl.RayWhite)
		rl.EndDrawing()
	}

}
