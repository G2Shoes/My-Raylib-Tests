package main

import (
	mrand "math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func CheckKeys() {

	var sentences [5]string

	sentences[0] = "boo!"
	sentences[1] = "fuck you blud!!"
	sentences[2] = "hi i teleported"
	sentences[3] = "fuckity fuckity fuck fuck fuck"
	sentences[4] = "Mr Garrison!"

	if rl.IsKeyPressed((rl.KeySpace)) {
		var goodint int = mrand.Intn(5)

		var sentenceint int = mrand.Intn(len(sentences))

		DrawTextAnchored(sentences[sentenceint], goodint, 25, rl.RayWhite)
	}
}
