package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Kabocha struct {
  Pawn
}

func loadKabocha() {
  kab = new(Kabocha)
  kab.src = rl.NewRectangle(0, 0, 48, 48) 
  kab.dest = rl.NewRectangle(200, 200, 60, 60) 
  kab.speed = 3.8
  kab.sprite.path = "./assets/sprites/kabocha.png"
  kab.pos = rl.Vector2{}

  kab.load() 
}

func (k *Kabocha) move() {
	// TODO: vector normalization for diagonal movement
  var ySprite int
  if k.dir != 0 {
    if k.dir&movingUp != 0 { 
      k.dest.Y -= k.speed
      ySprite = 1
    } else if k.dir&movingDown != 0 {
      k.dest.Y += k.speed
      ySprite = 0 
    }

    if k.dir&movingRight != 0 { 
      k.dest.X += k.speed
      ySprite = 3
    } else if k.dir&movingLeft != 0 {
      k.dest.X -= k.speed
      ySprite = 2
    }

    if frameCount % 6 == 1 { k.spriteFrame++ }
    if k.spriteFrame > 3 { k.spriteFrame = 0 }
  } else {
    if frameCount % 30 == 1 { k.spriteFrame++ }
    if k.spriteFrame > 1 { k.spriteFrame = 0 }
  }

  k.src.X = k.src.Width * float32(k.spriteFrame) 
  k.src.Y = k.src.Height * float32(ySprite) 

	k.setDirection(movingIdle)
}
