package main

import (
	"fmt"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	DBG       bool
	debugRecW int32 = 300
)

func DEBUG() {
	dRECXY(0, 0, debugRecW, SCRH, CA(z.Maroon, 50))
	var x, y float32 = 10, 10
	dTXT1XY("MS.X "+fmt.Sprintf("%.0f", MS.X)+" MS.Y "+fmt.Sprintf("%.0f", MS.Y), x, y)
	y += F1H
	dTXT1XY("CAM.Zoom "+fmt.Sprint(CAM.Zoom)+" gridH "+fmt.Sprint(gridH), x, y)
	y += F1H
	dTXT1XY("delta "+fmt.Sprint(delta)+" gridH "+fmt.Sprint(gridH), x, y)
	y += F1H

	//dREC(R(0, 0, float32(SCRW), float32(SCRH)), CA(BLACK(), 200))
	//SCROLLY = dIMSHEETSCROLL(TILES, float32(debugRecW)+10, 10, 10, 2, SCROLLY)

}
