package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func PtsDist(v1 rl.Vector2, v2 rl.Vector2) float64 {
  x := math.Pow(float64(v2.X - v1.X), 2) 
  y := math.Pow(float64(v2.Y - v1.Y), 2) 

  return math.Sqrt(x + y)
}

func DisVec(v1 rl.Vector2, v2 rl.Vector2) rl.Vector2 {
  return rl.NewVector2(
      v2.X - v1.X,
      v2.Y - v1.Y,
    )
}
