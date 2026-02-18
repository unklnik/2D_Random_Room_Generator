package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	SCRW, SCRH, FPS int32 = 1920, 1080, 60
	CAM             z.Camera2D
	CNT             z.Vector2
	FRAMES          int32

	STARTED bool

	//UNITS
	U1 float32 = 32

	UQTR, UHALF, U8TH, U3RD = U1 / 4, U1 / 2, U1 / 8, U1 / 3

	U2, U3, U4, U5, U6, U7, U8, U9, U10, U11, U12, U13, U14, U15 = U1 * 2, U1 * 3, U1 * 4, U1 * 5, U1 * 6, U1 * 7, U1 * 8, U1 * 9, U1 * 10, U1 * 11, U1 * 12, U1 * 13, U1 * 14, U1 * 15
)

func INIT() {
	z.SetTargetFPS(FPS)
	z.SetWindowState(z.FlagMsaa4xHint | z.FlagWindowUndecorated)
	z.HideCursor()

	//z.ToggleFullscreen()

	CNT = V2(float32(SCRW/2), float32(SCRH/2))
	CAM.Target = CNT
	CAM.Offset = V2(float32(SCRW/2), float32(SCRH/2))
	CAM.Rotation = 0.0
	CAM.Zoom = 1

	gamecolorLIST = append(gamecolorLIST, GREENYELLOW(), PALEGREEN(), LAWNGREEN(), YELLOWGREEN(), AQUA(), CYAN(), TOMATO(), MEDIUMPURPLE(), SANDYBROWN(), GOLDENROD())

	gameColor = gamecolorLIST[RINT(0, len(gamecolorLIST))]
	gameColorDARK = SEPIADARK()
	gameColor2 = z.Magenta
	gameColor3 = z.Orange

	bgColor = CA(gameColor, 20)

	for {
		gameColor2 = gamecolorLIST[RINT(0, len(gamecolorLIST))]
		if gameColor2 != gameColor {
			break
		}
	}
	for {
		shadowColor = gamecolorLIST[RINT(0, len(gamecolorLIST))]
		if shadowColor != gameColor && shadowColor != gameColor2 {
			break
		}
	}
	shadowColor = CA(shadowColor, 200)

	SHADER1ON = true
	SHADER2ON = true
	SCANLINES = true

	DMSBLOK = true

	SCROLLSPD = 0.1

	gridH = (int(SCRH) / int(blokW)) - 2
	gridW = gridH + gridH/2

	w := float32(gridW) * blokW
	h := float32(gridH) * blokW

	x := CNT.X - w/2
	y := CNT.Y - h/2

	borderR = R(x, y, w, h)
	borderINNER = RECSMLR(borderR, U2)
	startREC = RECFROMCNT(CNT, U4, U4)
	startREC.X = borderINNER.X

	INITSHADERS()

	mCOLORS()
	mTXT()
	mIMGS()
	mOBJ()
	mLEV()
	mPL()

}
func RESTART() {
	floorT = BGIM[RINT(0, len(BGIM))]
	wallT = TILES[RINT(0, len(TILES))]
	innerT = INNERIM[RINT(0, len(INNERIM))]

	room = nil
	roomOBJ = nil

	gameColor = gamecolorLIST[RINT(0, len(gamecolorLIST))]
	bgColor = CA(gameColor, 20)
	for {
		gameColor2 = gamecolorLIST[RINT(0, len(gamecolorLIST))]
		if gameColor2 != gameColor {
			break
		}
	}
	for {
		shadowColor = gamecolorLIST[RINT(0, len(gamecolorLIST))]
		if shadowColor != gameColor && shadowColor != gameColor2 {
			break
		}
	}
	shadowColor = CA(shadowColor, 200)
	mLEV()
	mPL()
}
func EXIT() {

	z.UnloadRenderTexture(RNDRTEX)
	z.UnloadRenderTexture(RNDRTEX2)

	z.UnloadFont(F1)
	z.UnloadFont(F2)

	for i := range RUBIK.fon {
		z.UnloadFont(RUBIK.fon[i])
	}

	for i := range len(TEX) {
		z.UnloadTexture(TEX[i])
	}

}
