package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (

	//MOUSE
	MS, MSCAM z.Vector2
	MSBLOKNUM int

	MSIN, DMSBLOK, MSINMENU, MSL, MSR, MSM, MSLDOWN, MSRDOWN, MSMDOWN bool
)

func INP() {

	//INP PLAYER
	move := z.NewVector2(0, 0)
	if z.IsKeyDown(z.KeyRight) || z.IsKeyDown(z.KeyD) {
		move.X += 1
		pl.lr = false
	}
	if z.IsKeyDown(z.KeyLeft) || z.IsKeyDown(z.KeyA) {
		move.X -= 1
		pl.lr = true
	}
	if z.IsKeyDown(z.KeyUp) || z.IsKeyDown(z.KeyW) {
		move.Y -= 1
	}
	if z.IsKeyDown(z.KeyDown) || z.IsKeyDown(z.KeyS) {
		move.Y += 1
	}
	move = NORMALIZE(move)
	if move.X == 0 && move.Y == 0 {
		pl.moving = false
	} else {

		v2 := pl.cnt
		v2.X += (move.X * pl.spd) * delta
		if cPLMOVE(v2) {
			pl.moving = true
			pl.cnt.X += (move.X * pl.spd) * delta
			uPL()
		} else {
			v2 = pl.cnt
		}
		v2.Y += (move.Y * pl.spd) * delta
		if cPLMOVE(v2) {
			pl.moving = true
			pl.cnt.Y += (move.Y * pl.spd) * delta
			uPL()
		}

	}
	//ZOOM
	if z.IsKeyPressed(z.KeyLeftBracket) {
		switch CAM.Zoom {
		case 0.25:
			CAM.Zoom = 2
		case 0.5:
			CAM.Zoom = 0.25
		case 1:
			CAM.Zoom = 0.5
		case 1.5:
			CAM.Zoom = 1
		case 2:
			CAM.Zoom = 1.5
		}
	} else if z.IsKeyPressed(z.KeyRightBracket) {
		switch CAM.Zoom {
		case 0.25:
			CAM.Zoom = 0.5
		case 0.5:
			CAM.Zoom = 1
		case 1:
			CAM.Zoom = 1.5
		case 1.5:
			CAM.Zoom = 2
		case 2:
			CAM.Zoom = 0.25
		}
	}

	//CORE
	if z.IsKeyPressed(z.KeyF1) {
		DBG = !DBG
	}
	if z.IsKeyPressed(z.KeyF2) {
		RESTART()
	}
	if z.IsKeyPressed(z.KeyF3) {

	}
	if z.IsKeyPressed(z.KeyF4) {
		SCANLINES = !SCANLINES
		//SHADER1ON = !SHADER1ON
	}
	if z.IsKeyPressed(z.KeyF5) {

	}
	if z.IsKeyPressed(z.KeyF6) {

	}
	if z.IsKeyPressed(z.KeyF7) {
		colorsON = !colorsON
	}
	//MOUSE
	MS = z.GetMousePosition()
	MSCAM = z.GetScreenToWorld2D(MS, CAM)

	if z.IsMouseButtonPressed(z.MouseButtonLeft) {
		MSL = true
	} else {
		MSL = false
	}
	if z.IsMouseButtonPressed(z.MouseButtonRight) {
		MSR = true
	} else {
		MSR = false
	}
	if z.IsMouseButtonPressed(z.MouseButtonMiddle) {
		MSM = true
	} else {
		MSM = false
	}
	if z.IsMouseButtonDown(z.MouseButtonLeft) {
		MSLDOWN = true
	} else {
		MSLDOWN = false
	}
	if z.IsMouseButtonDown(z.MouseButtonRight) {
		MSRDOWN = true
	} else {
		MSRDOWN = false
	}
	if z.IsMouseButtonDown(z.MouseButtonMiddle) {
		MSMDOWN = true
	} else {
		MSMDOWN = false
	}

}
func cMSREC(r z.Rectangle) bool {
	return cPOINTREC(MS, r)
}
