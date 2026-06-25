package main

import rl "github.com/gen2brain/raylib-go/raylib"

var PlayerSpawned bool = false

func main() {
	rl.InitWindow(800, 640, "Level Test 2")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	var floor []Floor
	var sky []Sky
	var testBlock []TestBlock

	var player Player

	floorTex := InitTex("Sprites/floor.png")
	skyTex := InitTex("Sprites/sky.png")
	tBlocktex := InitTex("Sprites/t.png")

	defer rl.UnloadTexture(floorTex)
	defer rl.UnloadTexture(skyTex)
	defer rl.UnloadTexture(tBlocktex)

	InitLevelGenerator(&floor, &sky, &testBlock, &player)

	camera := rl.Camera2D{
		Target:   rl.NewVector2(player.Position.X, 0),
		Offset:   rl.NewVector2(float32(rl.GetScreenWidth())/2, 0),
		Zoom:     1.0,
		Rotation: 0.0,
	}

	for !rl.WindowShouldClose() {

		camera.Target = rl.NewVector2(player.Position.X, 0)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Blank)

		rl.BeginMode2D(camera)
		PaintLevel(sky, floor, floorTex, skyTex, testBlock, tBlocktex)
		DrawPlayer(player, &player)
		rl.EndMode2D()
		DrawTextAnchored("Level Test 2", 3, 25, rl.RayWhite)
		rl.EndDrawing()
	}
}

func InitTex(filepath string) rl.Texture2D {
	tex := rl.LoadTexture(filepath)
	return tex
}

func DrawPlayer(p Player, plr *Player) {
	if PlayerSpawned {
		rl.DrawRectangleV(p.Position, p.Size, p.Color)
		Move(plr)
	}
}
