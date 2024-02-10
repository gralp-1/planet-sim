package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// window stuff
const WIDTH = 800
const HEIGHT = 600
const TITLE = "planet sim"

// planet stuff
const PLANET_RAD float32 = 10.0

// debug stuff
const DEBUG = true

// colour stuff
var VEL_LINE_COL = rl.Green
var ACC_LINE_COL = rl.Red
var PLANET_COL = rl.Color{255, 255, 255, 255}

type Planet struct {
	Pos rl.Vector2
	Vel rl.Vector2
	Acc rl.Vector2
	Col rl.Color
}

func (p *Planet) updatePos(centrePos rl.Vector2) {
	dist := rl.Vector2Distance(p.Pos, centrePos)
	_ = dist
	p.Acc = rl.Vector2Subtract(centrePos, p.Pos)
	p.Vel = rl.Vector2Add(p.Vel, p.Acc)
	tempVel := rl.Vector2{X: p.Vel.X * rl.GetFrameTime(), Y: p.Vel.Y * rl.GetFrameTime()}
	p.Pos = rl.Vector2Add(tempVel, p.Pos)
}
func (p *Planet) drawPlanet() {
	rl.DrawCircleV(p.Pos, PLANET_RAD, p.Col)
	if DEBUG {
		rl.DrawLineV(p.Pos, rl.Vector2Add(p.Pos, p.Vel), VEL_LINE_COL)
		rl.DrawLineV(p.Pos, rl.Vector2Add(p.Pos, p.Acc), ACC_LINE_COL)
	}
}

func main() {
	// TODO: dynamic orbiters
	centre := rl.Vector2{WIDTH / 2, HEIGHT / 2}
	moon := Planet{
		Pos: rl.Vector2Add(centre, rl.Vector2{X: 100}),
		Vel: rl.Vector2{Y: 2000},
		Col: PLANET_COL,
	}
	planet := Planet{Pos: centre, Col: rl.Orange}

	rl.InitWindow(WIDTH, HEIGHT, TITLE)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// update
		moon.updatePos(planet.Pos)

		// draw
		moon.drawPlanet()
		planet.drawPlanet()
		if DEBUG {
			rl.DrawFPS(10, 10)
		}

		rl.EndDrawing()
	}

}
