package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	pl              PLAYER
	hitPAUSEDEFTIME int32 = FPS * 2
)

type PLAYER struct {
	lr, moving             bool
	cnt                    z.Vector2
	spd, siz, opqRW, opqRH float32
	r, rCOLLIS             z.Rectangle
	anmMOVE, anmIDL        ANIM
	collisV2               []z.Vector2
	blokNUM, zi, hp, hpMAX int
	hitPAUSE               int32
}

// MARK: MAKE
func mPL() {
	PLIM = mIMSHEETXY("res/player.png", 6, 1, 0, 0, 16, 16, 0, 0)
	pl.anmMOVE = mANIMIMSHEET(PLIM, 16, 6)
	pl.anmIDL = mANIMIMSHEET(PLIM, 6, 6)
	pl.spd = U1DIV(3)
	pl.siz = U2
	pl.cnt = CNT
	pl.cnt.X = borderR.X + U2
	pl.r = RECFROMCNT(pl.cnt, pl.siz, pl.siz)
	pl.rCOLLIS = RECSMLR(pl.r, UQTR)
	pl.collisV2 = mPLCOLLISV2(pl.cnt)
	pl.opqRH = U3
	pl.opqRW = U10
	pl.hpMAX = 5
	pl.hp = 3

	uPL()
}
func mPLCOLLISV2(cnt z.Vector2) []z.Vector2 {
	v1 := cnt
	v1.Y -= pl.r.Height / 2
	v2 := cnt
	v2.X += pl.r.Width / 2
	v3 := cnt
	v3.Y += pl.r.Height / 2
	v4 := cnt
	v4.X -= pl.r.Width / 2
	return []z.Vector2{v1, v2, v3, v4}
}

// MARK: UP
func uPL() {
	pl.r = RECFROMCNT(pl.cnt, pl.siz, pl.siz)
	pl.rCOLLIS = RECSMLR(pl.r, UQTR)
	pl.collisV2 = mPLCOLLISV2(pl.cnt)
	if pl.hitPAUSE > 0 {
		pl.hitPAUSE--
	}
}

// MARK: DRAW
func dPL() {
	//SHADOW
	dIM(ETCIM[3], RECSHIFTURDL(pl.r, pl.r.Height/2, 3), z.White)
	if pl.lr {
		if pl.moving {
			pl.anmMOVE = dANIM(pl.anmMOVE, pl.r)
		} else {
			pl.anmIDL = dANIM(pl.anmIDL, pl.r)
		}
	} else {
		if pl.moving {
			pl.anmMOVE = dANIMFLIP(pl.anmMOVE, pl.r)
		} else {
			pl.anmIDL = dANIMFLIP(pl.anmIDL, pl.r)
		}
	}

	if DBG {
		for i := range pl.collisV2 {
			z.DrawCircleV(pl.collisV2[i], 8, z.Blue)
		}

	}

}

// MARK: CHECK
func cPLMOVE(cnt z.Vector2) bool {
	canmove := true
	checkV2 := mPLCOLLISV2(cnt)
	for i := range checkV2 {
		if !z.CheckCollisionPointRec(checkV2[i], borderR) {
			canmove = false
			break
		}
	}
	if canmove {
		for i := range room {
			for j := range checkV2 {
				if z.CheckCollisionPointRec(checkV2[j], room[i].r) {
					canmove = false
					break
				}
			}
		}
	}
	return canmove
}
