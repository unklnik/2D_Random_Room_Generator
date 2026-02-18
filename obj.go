package main

import (
	"math/rand/v2"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	objLIST []OBJ
)

type OBJ struct {
	r                         z.Rectangle
	nm                        string
	im                        IM
	cnt                       z.Vector2
	numTYPE                   int
	rSPIKES                   []z.Rectangle
	spikesUD, growshrinkONOFF bool

	velx, vely, maxSPD, maxGROWSHRINK, oWIDTH float32
}

// MARK: DRAW
func dOBJ() {
	for i := range roomOBJ {
		switch roomOBJ[i].numTYPE {
		case 2: //2 TURRET GUN
			dIMSHADOW(roomOBJ[i].im, roomOBJ[i].r, gameColor2, CA(z.Black, 200), UQTR)
		case 1: //1 NINJA START
			roomOBJ[i].im = dIMROTATESSHADOW(roomOBJ[i].im, roomOBJ[i].r, gameColor2, CA(z.Black, 200), UHALF, 8)
			roomOBJ[i] = GROWSHRINKOBJ(roomOBJ[i])
		case 0: //0 MOVING SPIKE BLOK
			dIMSHADOW(ETCIM[5], roomOBJ[i].rSPIKES[0], ORANGERED(), CA(z.Black, 200), UQTR)
			dIMROSHADOW(ETCIM[5], roomOBJ[i].rSPIKES[1], ORANGERED(), CA(z.Black, 200), 90, UQTR)
			dIMROSHADOW(ETCIM[5], roomOBJ[i].rSPIKES[2], ORANGERED(), CA(z.Black, 200), 180, UQTR)
			dIMROSHADOW(ETCIM[5], roomOBJ[i].rSPIKES[3], ORANGERED(), CA(z.Black, 200), 270, UQTR)
			dIM(ETCIM[6], roomOBJ[i].r, gameColor2)

		}
		if DBG {
			dRECLINES(roomOBJ[i].r, 2, BLUE())
		}
	}
}

// MARK: UP
func uOBJ() {

	for i := range roomOBJ {
		switch roomOBJ[i].numTYPE {
		case 1: //NINJA STAR

		case 0: //MOVING SPIKE BLOK
			//UP SPIKES
			w := roomOBJ[i].rSPIKES[0].Width
			if roomOBJ[i].spikesUD {
				if roomOBJ[i].rSPIKES[0].Y < roomOBJ[i].r.Y {
					roomOBJ[i].rSPIKES[0].Y++
					roomOBJ[i].rSPIKES[1].X--
					roomOBJ[i].rSPIKES[2].Y--
					roomOBJ[i].rSPIKES[3].X++
				} else {
					roomOBJ[i].spikesUD = false
				}
			} else {
				if roomOBJ[i].rSPIKES[0].Y > roomOBJ[i].r.Y-w {
					roomOBJ[i].rSPIKES[0].Y--
					roomOBJ[i].rSPIKES[1].X++
					roomOBJ[i].rSPIKES[2].Y++
					roomOBJ[i].rSPIKES[3].X--
				} else {
					roomOBJ[i].spikesUD = true
				}
			}
		}

		//MOVE
		checkR := roomOBJ[i].r
		checkR.X += roomOBJ[i].velx * delta
		bounceX, bounceY := false, false
		if cRECMOVE(checkR) {
			roomOBJ[i].r.X += roomOBJ[i].velx * delta
			if roomOBJ[i].numTYPE == 0 { //MOVE SPIKE BLOK RECS
				for j := range roomOBJ[i].rSPIKES {
					roomOBJ[i].rSPIKES[j].X += roomOBJ[i].velx * delta
				}
			}
		} else {
			checkR = roomOBJ[i].r
			bounceX = true
		}
		checkR.Y += roomOBJ[i].vely * delta
		if cRECMOVE(checkR) {
			roomOBJ[i].r.Y += roomOBJ[i].vely * delta
			if roomOBJ[i].numTYPE == 0 { //MOVE SPIKE BLOK RECS
				for j := range roomOBJ[i].rSPIKES {
					roomOBJ[i].rSPIKES[j].Y += roomOBJ[i].vely * delta
				}
			}
		} else {
			bounceY = true
		}
		//BOUNCE
		if roomOBJ[i].r.X <= borderR.X || roomOBJ[i].r.X+roomOBJ[i].r.Width >= borderR.X+borderR.Width || bounceX {
			roomOBJ[i].velx = -roomOBJ[i].velx
			roomOBJ[i].velx += (rand.Float32()*2 - 1) * (roomOBJ[i].maxSPD * 0.3)
			if roomOBJ[i].velx > roomOBJ[i].maxSPD {
				roomOBJ[i].velx = roomOBJ[i].maxSPD
			}
			if roomOBJ[i].velx < -roomOBJ[i].maxSPD {
				roomOBJ[i].velx = -roomOBJ[i].maxSPD
			}
		}
		if roomOBJ[i].r.Y <= borderR.Y || roomOBJ[i].r.Y+roomOBJ[i].r.Height >= borderR.Y+borderR.Height || bounceY {
			roomOBJ[i].vely = -roomOBJ[i].vely
			roomOBJ[i].vely += (rand.Float32()*2 - 1) * (roomOBJ[i].maxSPD * 0.3)
			if roomOBJ[i].vely > roomOBJ[i].maxSPD {
				roomOBJ[i].vely = roomOBJ[i].maxSPD
			}
			if roomOBJ[i].vely < -roomOBJ[i].maxSPD {
				roomOBJ[i].vely = -roomOBJ[i].maxSPD
			}
		}

	}
	uOBJPLCOLLIS() //OBJ COLLIS
}
func uOBJPLCOLLIS() {
	for i := range len(roomOBJ) {
		switch roomOBJ[i].numTYPE {
		case 0: //MOVING SPIKE BLOK
			if pl.hitPAUSE == 0 {
				for j := range roomOBJ[i].rSPIKES {
					if cRECREC(roomOBJ[i].rSPIKES[j], pl.rCOLLIS) {
						pl.hitPAUSE = hitPAUSEDEFTIME
						pl.hp--
					}
				}
			}
		}
	}
}

// MARK: MAKE
func mOBJ() {

	siz := U2

	//0 MOVING SPIKE BLOK
	o := OBJ{}
	o.numTYPE = 0
	o.nm = "moving spike blok"
	o.r = RECFROMCNT(CNT, siz, siz)
	o.cnt = CNT
	for range 4 {
		o.rSPIKES = append(o.rSPIKES, RECFROMCNT(o.cnt, siz/2, siz/2))
	}
	o.maxSPD = 8
	o.velx = RF32(-o.maxSPD/2, o.maxSPD/2)
	o.vely = RF32(-o.maxSPD/2, o.maxSPD/2)
	objLIST = append(objLIST, o)

	//1 NINJA STAR
	siz = U1 + UHALF
	o = OBJ{}
	o.numTYPE = 1
	o.nm = "ninja star"
	o.im = ETCIM[8]
	o.r = RECFROMCNT(CNT, siz, siz)
	o.oWIDTH = o.r.Width
	o.maxGROWSHRINK = o.r.Width + SIZMULTIDIV(o.r.Width, 20, 14)
	o.cnt = CNT
	o.maxSPD = 8
	o.velx = RF32(o.maxSPD/4, o.maxSPD)
	o.vely = RF32(o.maxSPD/4, o.maxSPD)
	objLIST = append(objLIST, o)

	//2 TURRET GUN
	o = OBJ{}
	o.numTYPE = 2
	o.nm = "turret gun"
	o.im = ETCIM[9]
	o.r = RECFROMCNT(CNT, siz, siz)
	o.cnt = CNT
	objLIST = append(objLIST, o)

}

// MARK: CHECK
func cOBJSTART() {
	for i := range len(roomOBJ) {
		for j := range len(room) {
			if cRECREC(roomOBJ[i].r, room[j].r) {
				roomOBJ[i].r = fRECPOSITION(roomOBJ[i].r)
				if roomOBJ[i].r.Y < borderINNER.Y {
					roomOBJ[i].r.Y = borderINNER.Y
				}
				switch roomOBJ[i].numTYPE {
				case 0: //0 MOVING SPIKE BLOK
					roomOBJ[i] = uSPIKEBLOKS(roomOBJ[i])
				}
			}
		}
	}
}

// MARK: UTILS
func uSPIKEBLOKS(o OBJ) OBJ {
	cnt := RECCNT(o.r)
	for i := range o.rSPIKES {
		o.rSPIKES[i] = RECFROMCNT(cnt, o.rSPIKES[i].Width, o.rSPIKES[i].Height)
	}

	return o
}
func GROWSHRINKOBJ(o OBJ) OBJ {
	if o.growshrinkONOFF {
		if o.r.Width > o.oWIDTH {
			w, h := o.r.Width, o.r.Height
			w--
			h--
			o.r = RECFROMCNT(RECCNT(o.r), w, h)
		} else if o.r.Width <= o.oWIDTH {
			o.r.Width = o.oWIDTH
			o.growshrinkONOFF = false
		}
	} else {
		if o.r.Width < o.maxGROWSHRINK {
			w, h := o.r.Width, o.r.Height
			w++
			h++
			o.r = RECFROMCNT(RECCNT(o.r), w, h)
		} else if o.r.Width >= o.maxGROWSHRINK {
			o.r.Width = o.maxGROWSHRINK
			o.growshrinkONOFF = true
		}
	}
	return o
}
