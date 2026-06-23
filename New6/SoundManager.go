package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func PlayOneShot(file string, duration int32) {

	go func() {
		noise := rl.LoadSound(file)
		rl.PlaySound(noise)

		time.Sleep(time.Duration(duration) * time.Second)

		defer rl.UnloadSound(noise)
	}()
}

