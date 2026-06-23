package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var ticks int
var gameover bool = false

func main() {

	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.InitWindow(288, 512, "RayBird")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	flappySprite1 := rl.LoadImage("Sprites/yellowbird-upflap.png")
	flappySprite2 := rl.LoadImage("Sprites/yellowbird-midflap.png")
	flappySprite3 := rl.LoadImage("Sprites/yellowbird-downflap.png")
	pipe := rl.LoadImage("Sprites/pipe-green.png")
	DayBg := rl.LoadImage("Sprites/background-day.png")
	flappySprite1Tex := InitTextues(flappySprite1)
	flappySprite2Tex := InitTextues(flappySprite2)
	flappySprite3Tex := InitTextues(flappySprite3)
	PipeSprite := InitTextues(pipe)
	Backdrop := InitTextues(DayBg)

	defer rl.UnloadTexture(Backdrop)
	defer rl.UnloadTexture(PipeSprite)
	defer rl.UnloadTexture(flappySprite1Tex)
	defer rl.UnloadTexture(flappySprite2Tex)
	defer rl.UnloadTexture(flappySprite3Tex)

	birdSprites := []rl.Texture2D{flappySprite1Tex, flappySprite2Tex, flappySprite3Tex}

	var pipes []Pipes

	var frame int

	var pipeFrame int

	birdie := Bird{
		Xpos:     30,
		Ypos:     float32(rl.GetScreenHeight()) / 2,
		Velocity: 2,
		Texture:  birdSprites[frame],
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if !gameover {
			rectBird := rl.NewRectangle(birdie.Xpos, birdie.Ypos, float32(birdie.Texture.Width), float32(birdie.Texture.Height))
			Update(birdSprites, &frame, &birdie, &pipes, PipeSprite, &pipeFrame, rectBird)

			rl.ClearBackground(rl.RayWhite)

			rl.DrawTexture(Backdrop, 0, 0, rl.RayWhite)
			DrawThePipes(pipes)

			PipeMove(pipes)
			DrawTheBird(birdie)
			DrawTextAnchored("Raybird Demo", 3, 15, rl.RayWhite)
		}

		if gameover {
			rl.ClearBackground(rl.Black)
			DrawTextAnchored("GAME OVER", 2, 45, rl.Red)
		}

		rl.EndDrawing()
	}
}

func InitTextues(img *rl.Image) rl.Texture2D {
	tex := rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)

	return tex
}

func Update(bird []rl.Texture2D, frame *int, birdo *Bird, pipe *[]Pipes, pipeTex rl.Texture2D, pipeFrameCount *int, birdRect rl.Rectangle) {
	ticks++
	(*pipeFrameCount)++

	if ticks >= 5 {
		*(frame)++

		if *frame >= len(bird) {
			*frame = 0
		}

		ticks = 0
		birdo.Texture = bird[*frame]
	}

	BirdUpdate(birdo)

	minTop := 80
	maxTop := 150

	minBottom := 200
	maxBottom := 400

	if *pipeFrameCount >= 120 {
		newPipe := Pipes{
			TopXPos:    int32(rl.GetScreenWidth()) + 30,
			BottomYpos: int32(minBottom + rand.N(maxBottom-minBottom+1)),
			TopYPos:    int32(minTop + rand.N(maxTop-minTop+1)),
			Sprite:     pipeTex,
			Active:     true,
		}

		*pipe = append(*pipe, newPipe)

		*pipeFrameCount = 0
	}

	for _, pussy := range *pipe {
		if pussy.Active {

			flippedX := float32(pussy.TopXPos) + float32(pussy.Sprite.Width)

			rectPipe1 := rl.NewRectangle(float32(pussy.TopXPos), float32(pussy.BottomYpos), float32(pussy.Sprite.Width)*1.5, float32(pussy.Sprite.Height))
			rectPipe2 := rl.NewRectangle(flippedX, float32(pussy.TopYPos)-float32(pussy.Sprite.Height), float32(pussy.Sprite.Width)*1.5, float32(pussy.Sprite.Height))

			if CheckCollisionBirdPipe(birdRect, rectPipe1, 1, rectPipe2) || CheckCollisionBirdPipe(birdRect, rectPipe1, 2, rectPipe2) {
				gameover = true
			}
		}
	}
}

func CheckCollisionBirdPipe(b rl.Rectangle, p rl.Rectangle, id int, p2 rl.Rectangle) bool {
	switch id {
	case 1:
		return rl.CheckCollisionRecs(b, p)
	case 2:
		return rl.CheckCollisionRecs(b, p2)
	}
	return false
}
