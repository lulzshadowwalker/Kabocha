package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Config struct {
	WindowHeight int
	WindowWidth  int
}

type Kabocha struct {
  sprite rl.Texture2D
  pos rl.Vector2
  src rl.Rectangle
  dest rl.Rectangle
  dir MovementDirection
  speed float32
  spriteFrame int
}

func loadKabocha() {
  kab = new(Kabocha)
  kab.load() 

  kab.src = rl.NewRectangle(0, 0, 48, 48) 
  kab.dest = rl.NewRectangle(200, 200, 60, 60) 
  kab.speed = 3.8

  kab.pos = rl.NewVector2(50, 50)
}

func (k *Kabocha) load() {
  k.sprite = rl.LoadTexture("./assets/sprites/kabocha.png")
}

func (k *Kabocha) unload() {
  rl.UnloadTexture(k.sprite)
}

func (k *Kabocha) render() {
  rl.DrawTexturePro(k.sprite, k.src, k.dest, k.pos, 0, rl.White)
}

func (k *Kabocha) update() {
  k.move()
}

func (k *Kabocha) setDirection(dir MovementDirection) {
	if dir == movingIdle {
		k.dir = 0
		return
	}

	k.dir = k.dir | dir
}

func (k *Kabocha) move() {
	// TODO, vector normalization for diagonal movement
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


type MovementDirection int

const (
	movingIdle      MovementDirection = 0b0
	movingDown      MovementDirection = 0b0001
	movingUp        MovementDirection = 0b0010
	movingLeft      MovementDirection = 0b0100
	movingRight     MovementDirection = 0b1000
)

var conf = Config{
  WindowHeight: rl.GetScreenHeight(),
  WindowWidth: rl.GetScreenWidth(),
}

var (
	running = true
  frameCount = 0 

  kab *Kabocha
)

func input() {
	if rl.IsKeyDown(rl.KeyW) {
		kab.setDirection(movingUp)
	} else if rl.IsKeyDown(rl.KeyS) {
		kab.setDirection(movingDown)
	}

	if rl.IsKeyDown(rl.KeyA) {
		kab.setDirection(movingLeft)
	} else if rl.IsKeyDown(rl.KeyD) {
		kab.setDirection(movingRight)
	}
}

func update() {
	running = !rl.WindowShouldClose()
  kab.update()
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

  kab.render() 

	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.RayWhite)

	rl.EndDrawing()
  frameCount++
}

var sample rl.Texture2D

func init() {
	rl.InitWindow(int32(conf.WindowWidth), int32(conf.WindowHeight), "Kabocha")
	rl.SetTargetFPS(60)

  sample = rl.LoadTexture("./assets/images/sample.jpg")
  loadKabocha()
}

func main() {
	defer quit()

	for running {
		input()
		update()
		render()
	}
}

func quit() {
	rl.CloseWindow()

  rl.UnloadTexture(sample)
}
