package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	MN                 []MENU
	MNSPD              = U1 / 2
	MENUISOPEN         bool
	MENUOPENNUM        int
	SCROLLY, SCROLLSPD float32
)

type MENU struct {
	r, rTAB    z.Rectangle
	open, msin bool
	c, cBG     z.Color
	nm         string
}

// MARK: DRAW
func dUI() {
	//PL HP
	siz := U1 + UHALF
	x, y := UQTR, UQTR
	for range pl.hpMAX {
		dIM(ETCIM[7], R(x, y, siz, siz), CA(LIGHTSLATEGRAY(), 100))
		y += siz
	}
	x, y = U8TH, U8TH
	for range pl.hp {
		dIM(ETCIM[7], R(x, y, siz, siz), RANRED())
		y += siz
	}

}

// MARK: BUTTONS
func ONOFFSWITCH(onoff bool, x, y, siz, offset float32, c z.Color) bool {
	r := R(x, y, siz, siz)
	r2 := R(x+offset, y+offset, siz-offset*2, siz-offset*2)
	r3 := R(x-2, y-2, siz+4, siz+4)
	c2 := CA(c, 100)
	dRECLINES(r, 2, c)
	if onoff {
		dREC(r2, c)
	} else {
		dREC(r2, c2)
	}
	if cMSREC(r) {
		dRECLINES(r3, 2, CA(c, fadeA))
		if MSL {
			onoff = !onoff
		}
	}
	return onoff
}

// MARK: SCROLL
func dSCROLLVERT(yChange, siz float32, c z.Color) float32 {
	r := R(float32(SCRW)-siz, 0, siz, siz)
	dIM(ETCIM[1], r, c)
	if cMSREC(r) && MSLDOWN {
		yChange += SCROLLSPD
	}
	r.Y += float32(SCRH) - siz
	dIM(ETCIM[0], r, c)
	if cMSREC(r) && MSLDOWN {
		yChange -= SCROLLSPD
	}
	return yChange
}
