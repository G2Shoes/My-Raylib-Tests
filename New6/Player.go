package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Xpos        int32
	Ypos        int32
	Speed       int32
	ColorTop    rl.Color
	ColorBottom rl.Color
	Size        float32
}

func DrawPlayer(p Player) {
	rl.DrawRectangleGradientV(p.Xpos, p.Ypos, int32(p.Size), int32(p.Size), p.ColorTop, p.ColorBottom)
}

func Move(p *Player) {

	sw := rl.GetScreenWidth()

	if rl.IsKeyDown(rl.KeyA) {
		p.Xpos -= p.Speed
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.Xpos += p.Speed
	}

	// Checking if player is outside the bounds
	if p.Xpos < 0 {
		PlayOneShot("snd/pop.wav", 1)
		p.Xpos = int32(p.Size)
	}

	if p.Xpos > int32(sw) {
		PlayOneShot("snd/pop.wav", 1)
		p.Xpos = int32(sw) - int32(p.Size*2)
	}
}

func Fire(p Player, pellet *[]Pellet) {
	if rl.IsKeyPressed(rl.KeySpace) {
		pelleting := Pellet{
			Speed:  15,
			Radius: 5,
			Xpos:   p.Xpos + (int32(p.Size) / 2),
			Ypos:   p.Ypos,
			Color:  rl.Yellow,
		}

		*pellet = append(*pellet, pelleting)
		PlayOneShot("snd/fire.wav", 1)
	}
}
