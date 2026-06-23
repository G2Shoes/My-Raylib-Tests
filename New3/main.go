package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	sw := 800
	sh := 600
	rl.InitWindow(int32(sw), int32(sh), "3D Prototype")
	defer rl.CloseWindow()

	camera := rl.Camera3D{
		Position:   rl.NewVector3(0, 5, 100),
		Target:     rl.NewVector3(0, 0, 0),
		Fovy:       45.0,
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Projection: rl.CameraPerspective,
	}

	rl.SetTargetFPS(60)
	rl.HideCursor()
	rl.DisableCursor()

	cameraMode := rl.CameraFree

	img := rl.LoadImage("imgs/grass.png")
	defer rl.UnloadImage(img)
	texGrass := rl.LoadTextureFromImage(img)
	defer rl.UnloadTexture(texGrass)

	mesh := rl.GenMeshCube(2.0, 2.0, 2.0)
	model := rl.LoadModelFromMesh(mesh)

	rl.SetMaterialTexture(model.Materials, rl.MapDiffuse, texGrass)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, cameraMode)
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)

		rl.BeginMode3D(camera)
		GenCubes(2.0, model)
		rl.EndMode3D()

		rl.DrawText("3D Testing App By Me", 10, 20, 25, rl.Black)

		rl.EndDrawing()
	}
}
