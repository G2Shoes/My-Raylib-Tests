package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Wall struct {
	position rl.Vector2
}

type Floor struct {
	position rl.Vector2
}

func Init() {
	fmt.Println("Initialized")
}
