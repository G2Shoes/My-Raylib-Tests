package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {

	rl.InitWindow(800, 640, "Camera 2D  Test")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	sw := rl.GetScreenWidth()
	sh := rl.GetScreenHeight()

	texturePlr := InitPlayerTex()
	textureObj := InitObjectTex()

	defer rl.UnloadTexture(texturePlr)
	defer rl.UnloadTexture(textureObj)

	positionObj := rl.NewVector2(200, 180)

	plr := Player{
		Speed:   5,
		Xpos:    0,
		Ypos:    int32((sh / 2) - 16),
		Texture: texturePlr,
	}

	camera := rl.Camera2D{
		Offset:   rl.NewVector2(float32((sw/2)-int(plr.Texture.Width)), float32((sh/2)-int(plr.Texture.Height))),
		Target:   rl.NewVector2(float32(plr.Xpos), float32(plr.Ypos)),
		Rotation: 0.0,
		Zoom:     1.5,
	}

	for !rl.WindowShouldClose() {
		Update(&plr)
		rl.BeginDrawing()

		camera.Target = rl.NewVector2(float32(plr.Xpos), float32(plr.Ypos))

		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.DarkGreen)
		rl.DrawTextureEx(textureObj, positionObj, 0, 1.0, rl.RayWhite)
		DrawPlayer(plr)
		rl.EndMode2D()

		DrawTextAnchored("Camera 2D Following Test", 0, 25, rl.White)

		rl.EndDrawing()
	}
}

func Update(p *Player) {
	move(p)
}

func InitPlayerTex() rl.Texture2D {
	plrSprite := rl.LoadImage("Textures/player.png")
	tex := rl.LoadTextureFromImage(plrSprite)
	rl.UnloadImage(plrSprite)

	return tex
}

func InitObjectTex() rl.Texture2D {
	objSprite := rl.LoadImage("Textures/obj.png")
	tex := rl.LoadTextureFromImage(objSprite)
	rl.UnloadImage(objSprite)

	return tex
}
