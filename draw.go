package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

func dCAM() {

	dFLOORWALLS()
	dINNER()

	dPL()

	if len(roomOBJ) > 0 {
		dOBJ()
	}

	dUI()

	if DBG {
		dRECLINES(borderR, 2, z.Orange)
		dRECLINES(borderINNER, 2, z.Orange)
		dRECLINES(startREC, 2, z.Orange)
		dRECLINES(pl.r, 2, BLUE())
		dRECLINES(pl.rCOLLIS, 2, WHITE())
	}

}

func dNOCAM() {

	if SCANLINES {
		dSCAN(5, 2, CA(z.Black, RUINT8(20, 50)))
	}

	if DBG {
		DEBUG()
	}
	if colorsON {
		dCOLORS()
	}

	//MS CURSOR
	dIM(ETCIM[2], R(MS.X, MS.Y, U1, U1), gameColor2)

}
