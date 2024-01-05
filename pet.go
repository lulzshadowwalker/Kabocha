package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Pet struct {
  Pawn
}

func loadChick() {
  chick = new(Pet)
  chick.src = rl.NewRectangle(0, 0, 16, 16) 
  chick.dest = rl.NewRectangle(200, 200, 60, 60) 
  chick.speed = 3.8
  chick.sprite.path = "./assets/sprites/chick.png"
  chick.pos = rl.Vector2{}

  chick.load()
}
