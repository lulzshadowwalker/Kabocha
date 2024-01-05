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
  chick.dest = rl.NewRectangle(0, 0, 40, 40) 
  chick.speed = 1.5
  chick.sprite.path = "./assets/sprites/chick.png"
  chick.pos = rl.Vector2{}

  chick.load()
}

func (p *Pet) update() {
  p.move()
}

func (p *Pet) move() {
  vec := DisVec(p.pos, kab.pos)
  dist := PtsDist(p.pos, kab.pos)

  var ySprite int
  if dist > 15 {
    ySprite = 1

    if vec.X > 0 {
      p.pos.X += p.speed
    } else {
      p.pos.X -= p.speed 
    }

    if vec.Y > 0 {
      p.pos.Y += p.speed
    } else {
      p.pos.Y -= p.speed 
    }

    if frameCount % 15 == 0 { p.spriteFrame++ } 
    if p.spriteFrame > 3 { p.spriteFrame = 0 }
  } else {
    ySprite = 0 

    if frameCount % 35 == 0 { p.spriteFrame++ }
    if p.spriteFrame > 1 { p.spriteFrame = 0 }
  }

  p.src.X = p.src.Width * float32(p.spriteFrame) 
  p.src.Y = p.src.Height * float32(ySprite) 

  p.setDirection(movingIdle)
}
