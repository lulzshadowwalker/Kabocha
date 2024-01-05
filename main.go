package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	running = true
  frameCount = 0 

  kab *Kabocha
  chick *Pet
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
  chick.update()
}

func render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

  kab.render() 
  chick.render()

	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.RayWhite)

	rl.EndDrawing()
  frameCount++
}

func init() {
	rl.InitWindow(int32(conf.WindowWidth), int32(conf.WindowHeight), "Kabocha")
	rl.SetTargetFPS(60)

  loadKabocha()
  loadChick() 
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
}
