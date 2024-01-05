package main 

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	movingIdle      MovementDirection = 0b0
	movingDown      MovementDirection = 0b0001
	movingUp        MovementDirection = 0b0010
	movingLeft      MovementDirection = 0b0100
	movingRight     MovementDirection = 0b1000
)

type MovementDirection int

// TODO: properly update Pawn.pos
type Pawn struct {
  sprite KSprite
  pos rl.Vector2
  src rl.Rectangle
  dest rl.Rectangle
  dir MovementDirection
  speed float32
  spriteFrame int
}

type KSprite struct {
  texture rl.Texture2D 
  path string
}

func (p *Pawn) load() {
  p.sprite.texture = rl.LoadTexture(p.sprite.path)
}

func (p *Pawn) unload() {
  rl.UnloadTexture(p.sprite.texture)
}

func (p *Pawn) render() {
  rl.DrawTexturePro(p.sprite.texture, p.src, p.dest, p.pos, 0, rl.White)
}

func (k *Kabocha) update() {
  k.move()
}

func (p *Pawn) setDirection(dir MovementDirection) {
	if dir == movingIdle {
		p.dir = 0
		return
	}

	p.dir = p.dir | dir
}

