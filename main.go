package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// window stuff
const WIDTH = 1512
const HEIGHT = 830
const TITLE = "planet sim"

// planet stuff
const PLANET_RAD float32 = 10.0

// flags
const DEBUG = true
const DRAW_ORBIT = false

// colour stuff (can't make this constant but we can pretend)
var VEL_LINE_COL = rl.Green
var ACC_LINE_COL = rl.Red
var PLANET_COL = rl.Color{R: 255, G: 255, B: 255, A: 255}

type Planet struct {
	Pos rl.Vector2
	Vel rl.Vector2
	Acc rl.Vector2
	Col rl.Color
}

func (p *Planet) updatePos(centrePos rl.Vector2) {
	/* planet movement logic, based off of Verlet integration but pretty much just:
	new acceleration = vector to orbit planet
	new velocity = velocity + new acceleration
	new pos = pos + new velocity (adjusted for frame time)
	*/
	// dist := rl.Vector2Distance(p.Pos, centrePos)
	p.Acc = rl.Vector2Subtract(centrePos, p.Pos)
	p.Vel = rl.Vector2Add(p.Vel, p.Acc)

	// adjust for fps so it's always smooth
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
	// TODO: flags (debug, draw orbit etc) passed in as flags rather than constant, maybe at compile time?
	// TODO: dynamic orbiters
	centre := rl.Vector2{X: WIDTH / 2, Y: HEIGHT / 2}

	// the one that moves
	moon := Planet{
		Pos: rl.Vector2Add(centre, rl.Vector2{X: 100}),
		Vel: rl.Vector2{Y: 100},
		Col: PLANET_COL,
	}
	moon2 := Planet{
		Pos: rl.Vector2Add(centre, rl.Vector2{X: 400}),
		Vel: rl.Vector2{Y: 400},
		Col: rl.Gray,
	}
	planet := Planet{Pos: centre, Col: rl.DarkGreen}

	rl.InitWindow(WIDTH, HEIGHT, TITLE)
	defer rl.CloseWindow()
	rl.SetTargetFPS(30)

	// used to store the points that the orbiter has already been at
	// using map[pos]bool instead of list of positions as a hacky way to prevent storing tonnes of duplicate values
	orbitPoints := map[rl.Vector2]bool{}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		moon.updatePos(planet.Pos)
		moon2.updatePos(planet.Pos)

		// draw
		moon.drawPlanet()
		planet.drawPlanet()
		moon2.drawPlanet()

		// draw orbits
		if DRAW_ORBIT {
			orbitPoints[moon.Pos] = true
			for pos := range orbitPoints {
				rl.DrawCircleV(pos, 2, rl.RayWhite)
			}
		}

		if DEBUG {
			rl.DrawFPS(10, 10)
			rl.DrawText(fmt.Sprintf("Pos: %v", moon.Pos), 10, 25, 15, rl.DarkGreen)
			rl.DrawText(fmt.Sprintf("Vel: %v", moon.Vel), 10, 40, 15, rl.DarkGreen)
			rl.DrawText(fmt.Sprintf("Acc: %v", moon.Acc), 10, 55, 15, rl.DarkGreen)
			rl.DrawText(fmt.Sprintf("Spd: %v", rl.Vector2Length(moon.Acc)), 10, 70, 15, rl.DarkGreen)
			rl.DrawText(fmt.Sprintf("orbit points: %v", len(orbitPoints)), 10, 85, 15, rl.DarkGreen)
		}
		rl.EndDrawing()
	}
}
