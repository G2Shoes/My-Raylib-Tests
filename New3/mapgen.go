package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func GenCubes(gap float32, model rl.Model) {
	sizeX := 10
	sizeY := 5
	sizeZ := 10

	for y := 0; y < sizeY; y++ {
		for z := 0; z < sizeZ; z++ {
			for x := 0; x < sizeX; x++ {
				posX := float32(x) * gap
				posY := float32(y) * gap
				posZ := float32(z) * gap

				position := rl.NewVector3(posX, posY, posZ)
				rl.DrawModel(model, position, 1.0, rl.White)
			}
		}
	}
}
