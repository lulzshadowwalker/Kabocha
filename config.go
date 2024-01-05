package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var conf = Config {
  WindowHeight: rl.GetScreenHeight() / 2,
  WindowWidth: rl.GetScreenWidth() / 2,
}

type Config struct {
	WindowHeight int
	WindowWidth  int
}
