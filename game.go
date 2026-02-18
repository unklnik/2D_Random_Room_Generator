package main

import z "github.com/gen2brain/raylib-go/raylib"

var (
	gridH, gridW, levNUM, gridWNUM int
	blokW                          = U1
	borderR, borderINNER, startREC z.Rectangle

	selectTILE = BLANKNUM

	room []BLOK

	roomOBJ []OBJ
)

// MARK: MAKE
func mLEVEXTRAS() {

	choose := RINT(0, 11)
	choose = 0
	switch choose {
	case 2: //TURRET GUN
		o := objLIST[2]
		o.r = fRECUDBORDEROUTER(o.r)
		roomOBJ = append(roomOBJ, o)
	case 1: //NINJA STAR
		o := objLIST[1]
		o.r = fRECPOSITION(o.r)
		o = OBJSPD(o)
		roomOBJ = append(roomOBJ, o)
	case 0: //MOVING BLOK
		o := objLIST[0]
		o.r = fRECPOSITION(o.r)
		o = uSPIKEBLOKS(o)
		o = OBJSPD(o)
		roomOBJ = append(roomOBJ, o)
	}

}
func mLEV() {

	choose := RINT(0, 30)
	siz := U2
	osiz := siz
	spc := UHALF

	switch choose {
	case 29: //LONG REC RAND BLOKS
		b := BLOK{}
		w := RINT(7, 16)
		h := 2
		b.r = R(borderINNER.X+siz*2, borderINNER.Y+RF32(siz*2, siz*8), siz, siz)
		if FLIPCOIN() {
			b.r.X += siz
			if FLIPCOIN() {
				b.r.X += siz
			}
		}
		ox := b.r.X
		a := w * h
		oa := a
		c := 0
		for a > 0 {
			room = append(room, b)
			if a > oa-w {
				if ROLL6() > 4 {
					b2 := b
					b2.r.Y -= siz
					room = append(room, b2)
					if FLIPCOIN() {
						b2.r.Y -= siz
						room = append(room, b2)
					}
				}
			} else {
				if ROLL6() > 4 {
					b2 := b
					b2.r.Y += siz
					room = append(room, b2)
					if FLIPCOIN() {
						b2.r.Y += siz
						room = append(room, b2)
					}
				}
			}
			b.r.X += siz

			c++
			if c == w {
				b.r.X = ox
				b.r.Y += siz
				c = 0
			}
			a--
		}

	case 28: //TRIANGLES
		b := BLOK{}
		b.r = R(borderINNER.X+siz*6, borderINNER.Y+RF32(siz*5, siz*9), siz, siz)
		if FLIPCOIN() {
			b.r.X += siz
			if FLIPCOIN() {
				b.r.X += siz
			}
		}

		room = append(room, b)
		b.r.X -= siz
		b.r.Y += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		b.r.Y += siz
		room = append(room, b)
		b.r.X -= siz
		room = append(room, b)
		b.r.X -= siz
		room = append(room, b)
		b.r.X -= siz
		room = append(room, b)
		b.r.X -= siz
		room = append(room, b)
		b.r.X -= siz
		b.r.Y += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)

		b.r.X = CNT.X
		if FLIPCOIN() {
			b.r.X += siz
			if FLIPCOIN() {
				b.r.X += siz
			}
		}
		b.r.Y = borderINNER.Y
		b.r.Y += RF32(siz, siz*5)

		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X -= siz
		b.r.Y += siz
		room = append(room, b)
		b.r.X -= siz
		room = append(room, b)
		b.r.X -= siz
		room = append(room, b)
		b.r.X -= siz
		room = append(room, b)
		b.r.X -= siz
		room = append(room, b)
		b.r.X += siz
		b.r.Y += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X += siz
		room = append(room, b)
		b.r.X -= siz
		b.r.Y += siz
		room = append(room, b)

	case 27: //DIAMOND RAND BLOKS
		b := BLOK{}
		b.r = RECFROMCNT(CNT, siz, siz)
		b.r.Y = borderINNER.Y + siz/2
		if FLIPCOIN() {
			b.r.X += RF32(-siz*4, siz*4)
		}

		b.r.X -= siz
		b.r.Y += siz
		if ROLL6() != 6 {
			room = append(room, b)
		}
		b2 := b
		b2.r.X += siz * 2
		if ROLL6() != 6 {
			room = append(room, b2)
		}
		b.r.X -= siz
		b.r.Y += siz
		if ROLL6() != 6 {
			room = append(room, b)
		}
		b2 = b
		b2.r.X += siz * 4
		if ROLL6() != 6 {
			room = append(room, b2)
		}
		b.r.X -= siz
		b.r.Y += siz
		if ROLL6() != 6 {
			room = append(room, b)
		}
		b2 = b
		b2.r.X += siz * 6
		if ROLL6() != 6 {
			room = append(room, b2)
		}
		b.r.X -= siz
		b.r.Y += siz
		if ROLL6() != 6 {
			room = append(room, b)
		}
		b2 = b
		b2.r.X += siz * 8
		if ROLL6() != 6 {
			room = append(room, b2)
		}
		b.r.Y += siz * 4
		b2.r.Y += siz * 4

		if ROLL6() != 6 {
			room = append(room, b)
		}
		b2 = b
		b2.r.X += siz * 8
		if ROLL6() != 6 {
			room = append(room, b2)
		}
		b.r.X += siz
		b.r.Y += siz
		if ROLL6() != 6 {
			room = append(room, b)
		}
		b2 = b
		b2.r.X += siz * 6
		if ROLL6() != 6 {
			room = append(room, b2)
		}
		b.r.X += siz
		b.r.Y += siz
		if ROLL6() != 6 {
			room = append(room, b)
		}
		b2 = b
		b2.r.X += siz * 4
		if ROLL6() != 6 {
			room = append(room, b2)
		}
		b.r.X += siz
		b.r.Y += siz
		if ROLL6() != 6 {
			room = append(room, b)
		}
		b2 = b
		b2.r.X += siz * 2
		if ROLL6() != 6 {
			room = append(room, b2)
		}
	case 26: //DIAMOND
		b := BLOK{}
		b.r = RECFROMCNT(CNT, siz, siz)
		b.r.Y = borderINNER.Y + siz/2
		if FLIPCOIN() {
			b.r.X += RF32(-siz*4, siz*4)
		}

		b.r.X -= siz
		b.r.Y += siz
		room = append(room, b)
		b2 := b
		b2.r.X += siz * 2
		room = append(room, b2)
		b.r.X -= siz
		b.r.Y += siz
		room = append(room, b)
		b2 = b
		b2.r.X += siz * 4
		room = append(room, b2)
		b.r.X -= siz
		b.r.Y += siz
		room = append(room, b)
		b2 = b
		b2.r.X += siz * 6
		room = append(room, b2)
		b.r.X -= siz
		b.r.Y += siz
		room = append(room, b)
		b2 = b
		b2.r.X += siz * 8
		room = append(room, b2)
		b.r.Y += siz * 4
		b2.r.Y += siz * 4

		room = append(room, b)
		b2 = b
		b2.r.X += siz * 8
		room = append(room, b2)
		b.r.X += siz
		b.r.Y += siz
		room = append(room, b)
		b2 = b
		b2.r.X += siz * 6
		room = append(room, b2)
		b.r.X += siz
		b.r.Y += siz
		room = append(room, b)
		b2 = b
		b2.r.X += siz * 4
		room = append(room, b2)
		b.r.X += siz
		b.r.Y += siz
		room = append(room, b)
		b2 = b
		b2.r.X += siz * 2
		room = append(room, b2)

	case 25: //RANDOM SHORT LINES HORIZ VERT
		b := BLOK{}
		b.r = R(borderINNER.X+siz*2, borderINNER.Y, siz, siz)
		for b.r.X < borderINNER.X+borderINNER.Width-siz*3 {
			if FLIPCOIN() {
				b.r.Y += RF32(0, borderINNER.Height-siz*2)
				for range 3 {
					room = append(room, b)
					b.r.X += siz
					if b.r.X >= borderINNER.X+borderINNER.Width-siz*3 {
						break
					}
				}
				if FLIPCOIN() {
					room = append(room, b)
				}
			} else {
				b.r.Y += RF32(0, borderINNER.Height/2)
				for range 3 {
					room = append(room, b)
					b.r.Y += siz
				}
				if FLIPCOIN() {
					room = append(room, b)
				}
			}
			b.r.X += siz * 2
			b.r.Y = borderINNER.Y
		}

	case 24: //RANDOM BLOKS
		b := BLOK{}
		osiz := siz
		num := RINT(20, 31)
		for num > 0 {
			siz += RF32(-siz/4, siz)
			x := RF32(borderINNER.X, borderINNER.X+borderINNER.Width-siz*2)
			y := RF32(borderINNER.Y, borderINNER.Y+borderINNER.Height-siz)
			b.r = R(x, y, siz, siz)
			if cADDBLOKBORDER(b, siz) {
				room = append(room, b)
			}
			siz = osiz
			num--
		}
	case 23: //SINGLE LARGE BLOK
		siz += RF32(siz*2, siz*4)
		b := BLOK{}
		b.r = RECFROMCNT(CNT, siz, siz)
		if FLIPCOIN() {
			b.r.X += RF32(-siz, siz)
		}
		if FLIPCOIN() {
			b.r.Y += RF32(-siz, siz)
		}
		room = append(room, b)
	case 22: //ZIGZAG LARGE BLOKS
		siz += RF32(siz/2, siz*2)
		b := BLOK{}
		b.r = R(borderINNER.X+siz, borderINNER.Y, siz, siz)
		b.r.Y += RF32(0, siz*2)
		flip := false
		for b.r.X < borderINNER.X+borderINNER.Width-(siz*4) {
			if cADDBLOK(b) {
				room = append(room, b)
			}
			if flip {
				b.r.X += siz
				b.r.Y -= siz
				if cADDBLOK(b) {
					room = append(room, b)
				}
				b.r.X += siz
				b.r.Y -= siz
				if cADDBLOK(b) {
					room = append(room, b)
				}
				flip = false
			} else {
				b.r.X += siz
				b.r.Y += siz
				if cADDBLOK(b) {
					room = append(room, b)
				}
				b.r.X += siz
				b.r.Y += siz
				if cADDBLOK(b) {
					room = append(room, b)
				}
				flip = true
			}
		}
	case 21: //PARALLEL ZIGZAG MISSING BLOKS
		b := BLOK{}
		b.r = R(borderINNER.X+U2, borderINNER.Y+U2, siz, siz)
		flip := false
		for b.r.X < borderINNER.X+borderINNER.Width-(siz*4) {
			if ROLL6() < 5 {
				room = append(room, b)
			}
			if flip {
				b.r.X += siz
				b.r.Y -= siz
				if ROLL6() < 5 {
					room = append(room, b)
				}
				b.r.X += siz
				b.r.Y -= siz
				if ROLL6() < 5 {
					room = append(room, b)
				}
				flip = false
			} else {
				b.r.X += siz
				b.r.Y += siz
				if ROLL6() < 5 {
					room = append(room, b)
				}
				b.r.X += siz
				b.r.Y += siz
				if ROLL6() < 5 {
					room = append(room, b)
				}
				flip = true
			}
		}
		b.r = R(borderINNER.X+U2, borderINNER.Y+borderINNER.Height-siz*5, siz, siz)
		flip = false
		for b.r.X < borderINNER.X+borderINNER.Width-(siz*4) {
			if ROLL6() < 5 {
				room = append(room, b)
			}
			if flip {
				b.r.X += siz
				b.r.Y -= siz
				if ROLL6() < 5 {
					room = append(room, b)
				}
				b.r.X += siz
				b.r.Y -= siz
				if ROLL6() < 5 {
					room = append(room, b)
				}
				flip = false
			} else {
				b.r.X += siz
				b.r.Y += siz
				if ROLL6() < 5 {
					room = append(room, b)
				}
				b.r.X += siz
				b.r.Y += siz
				if ROLL6() < 5 {
					room = append(room, b)
				}
				flip = true
			}
		}
	case 20: //PARALLEL ZIGZAG
		b := BLOK{}
		b.r = R(borderINNER.X+U2, borderINNER.Y+U2, siz, siz)
		flip := false
		for b.r.X < borderINNER.X+borderINNER.Width-(siz*4) {
			room = append(room, b)
			if flip {
				b.r.X += siz
				b.r.Y -= siz
				room = append(room, b)
				b.r.X += siz
				b.r.Y -= siz
				room = append(room, b)
				flip = false
			} else {
				b.r.X += siz
				b.r.Y += siz
				room = append(room, b)
				b.r.X += siz
				b.r.Y += siz
				room = append(room, b)
				flip = true
			}
		}
		b.r = R(borderINNER.X+U2, borderINNER.Y+borderINNER.Height-siz*5, siz, siz)
		flip = false
		for b.r.X < borderINNER.X+borderINNER.Width-(siz*4) {
			room = append(room, b)
			if flip {
				b.r.X += siz
				b.r.Y -= siz
				room = append(room, b)
				b.r.X += siz
				b.r.Y -= siz
				room = append(room, b)
				flip = false
			} else {
				b.r.X += siz
				b.r.Y += siz
				room = append(room, b)
				b.r.X += siz
				b.r.Y += siz
				room = append(room, b)
				flip = true
			}
		}
	case 19: // PARALLEL BLOCKS RAND SIZ
		siz += RF32(siz/2, siz*2)
		b := BLOK{}
		b.r = R(borderINNER.X+U2, borderINNER.Y+U2, siz, siz)
		for b.r.X < borderINNER.X+borderINNER.Width-(siz*2) {
			if FLIPCOIN() {
				change := RF32(-siz/4, siz/4)
				b.r.Width += change
				b.r.Height += change
			}
			room = append(room, b)
			b2 := b
			b2.r.Y += borderINNER.Height - siz*2
			if FLIPCOIN() {
				change := RF32(-siz/4, siz/4)
				b2.r.Width += change
				b2.r.Height += change
			}
			room = append(room, b2)
			b.r.X += siz * 2
		}
	case 18: // PARALLEL BLOCKS
		siz += RF32(siz, siz*3)
		b := BLOK{}
		b.r = R(borderINNER.X+U2*2, borderINNER.Y+U2, siz, siz)
		for b.r.X < borderINNER.X+borderINNER.Width-(siz+spc*2) {
			room = append(room, b)
			b2 := b
			b2.r.Y += borderINNER.Height - siz*2
			room = append(room, b2)
			b.r.X += siz * 2
		}
	case 17: //LONG VERT U
		b := BLOK{}
		xChange := float32(RINT(4, 8)) * siz
		b.r = R(borderINNER.X+xChange, borderINNER.Y+siz, siz, siz)
		//ox := b.r.X
		oy := b.r.Y
		yChange := float32(RINT(1, 5))
		gapChange := RINT(4, 8)
		gap := float32(gapChange) * siz
		for b.r.Y < borderINNER.Y+borderINNER.Height-siz*yChange {
			if ROLL6() < 5 {
				room = append(room, b)
			}
			b2 := b
			b2.r.X += gap
			if ROLL6() < 5 {
				room = append(room, b2)
			}
			b.r.Y += siz
		}
		b.r.Y -= siz
		for gapChange > 0 {
			b.r.X += siz
			if cADDBLOK(b) {
				room = append(room, b)
			}
			gapChange--
		}
		b.r.Y = oy
		yChange = float32(RINT(1, 5))
		gapChange = RINT(3, 6)
		for gapChange > 0 {
			b.r.X += siz
			if cADDBLOK(b) {
				room = append(room, b)
			}
			gapChange--
		}
		gap = float32(gapChange) * siz
		for b.r.Y < borderINNER.Y+borderINNER.Height-siz*yChange {
			if ROLL6() < 5 {
				room = append(room, b)
			}
			b.r.Y += siz
		}

	case 16: //CORNER L
		b := BLOK{}
		num := 3
		onum := num
		b.r.X = borderINNER.X + siz*5
		b.r.Y = borderINNER.Y
		ox := b.r.X
		oy := b.r.Y
		b.r.Width, b.r.Height = siz, siz
		for num > 0 {
			b.r.Y += siz
			room = append(room, b)
			num--
		}
		num = onum
		num--
		b.r.X -= siz
		for num > 0 {
			room = append(room, b)
			b.r.X -= siz
			num--
		}

		b.r.X = ox
		b.r.Y = oy + siz/2
		oy = b.r.Y
		b.r.X += (siz + spc) * 2
		ox = b.r.X
		num = onum
		num += 2
		onum = num
		for num > 0 {
			b.r.Y += siz
			room = append(room, b)
			num--
		}
		num = onum
		num--
		b.r.X -= siz
		for num > 0 {
			room = append(room, b)
			b.r.X -= siz
			num--
		}

		b.r.X = ox
		b.r.Y = oy + siz/2
		oy = b.r.Y
		b.r.X += (siz + spc) * 2
		ox = b.r.X
		num = onum
		num += 2
		onum = num
		for num > 0 {
			b.r.Y += siz
			room = append(room, b)
			num--
		}
		num = onum
		num--
		b.r.X -= siz
		for num > 0 {
			room = append(room, b)
			b.r.X -= siz
			num--
		}

		b.r.X = ox
		b.r.Y = oy + siz/2
		oy = b.r.Y
		b.r.X += (siz + spc) * 2
		ox = b.r.X
		num = onum
		num += 2
		onum = num
		for num > 0 {
			b.r.Y += siz
			room = append(room, b)
			num--
		}
		num = onum
		num--
		b.r.X -= siz
		for num > 0 {
			room = append(room, b)
			b.r.X -= siz
			num--
		}

	case 15: //LRG BLOK CROSS
		b := BLOK{}
		w := RF32(siz*2, siz*4)
		b.r = RECFROMCNT(CNT, w, w)
		room = append(room, b)
		b2 := b
		b2.r.X -= w + spc
		b2.r.Y -= w
		wChange := RF32(siz/2, siz)
		b2.r.Width -= wChange
		b2.r.Height -= wChange
		if FLIPCOIN() {
			b2.r.X -= RF32(siz, siz*2)
		}
		room = append(room, b2)

		b2 = b
		b2.r.X += w + spc
		b2.r.Y -= w
		wChange = RF32(siz/2, siz)
		b2.r.Width -= wChange
		b2.r.Height -= wChange
		if FLIPCOIN() {
			b2.r.X += RF32(siz, siz*2)
		}
		room = append(room, b2)

		b2 = b
		b2.r.X += w + spc
		b2.r.Y += w
		wChange = RF32(siz/2, siz)
		b2.r.Width -= wChange
		b2.r.Height -= wChange
		if FLIPCOIN() {
			b2.r.X += RF32(siz, siz*2)
		}
		room = append(room, b2)

		b2 = b
		b2.r.X -= w + spc
		b2.r.Y += w
		wChange = RF32(siz/2, siz)
		b2.r.Width -= wChange
		b2.r.Height -= wChange
		if FLIPCOIN() {
			b2.r.X -= RF32(siz, siz*2)
		}
		room = append(room, b2)

	case 14: //CENTER CROSS RAND HORIZ
		b := BLOK{}
		b.r = RECFROMCNT(CNT, siz, siz)
		b.r.Y = borderINNER.Y + siz/2
		for b.r.Y+siz < borderINNER.Y+borderINNER.Height {
			if ROLL6() != 6 {
				room = append(room, b)
				b.r.Y += siz + spc
			} else {
				b.r.Y += siz + siz/2
			}
		}
		b.r.X = borderINNER.X + siz*3
		if FLIPCOIN() {
			b.r.X += siz
			if FLIPCOIN() {
				b.r.X += siz
			}
		}
		b.r.Y = borderINNER.Y + siz/2
		b.r.Y += RF32(0, borderINNER.Height-siz*4)
		for b.r.X+siz < borderINNER.X+borderINNER.Width-siz*2 {
			if ROLL6() != 6 {
				if cADDBLOKBORDER(b, U8TH) {
					room = append(room, b)
				}
				b.r.X += siz + spc
			} else {
				b.r.X += siz + siz/2
			}
		}
	case 13: //CENTER DIAG CROSS
		b := BLOK{}
		b.r = RECFROMCNT(CNT, siz, siz)
		b2, b3, b4 := b, b, b
		num := RINT(2, 5)
		for num > 0 {
			b.r.X -= siz + spc
			b.r.Y -= siz + spc
			room = append(room, b)
			b2.r.X += siz + spc
			b2.r.Y -= siz + spc
			room = append(room, b2)
			b3.r.X -= siz + spc
			b3.r.Y += siz + spc
			room = append(room, b3)
			b4.r.X += siz + spc
			b4.r.Y += siz + spc
			room = append(room, b4)
			num--
		}
	case 12: //RAND LINE FROM RIGHT RAND SIZ
		osiz := siz
		b := BLOK{}
		siz = RF32(siz, siz*3)
		sizPREV := siz
		b.r = R(borderINNER.X+borderINNER.Width-siz*2, borderINNER.Y+siz, siz, siz)
		room = append(room, b)
		siz = osiz
		num := RINT(8, 15)
		for num > 0 {
			siz = RF32(siz, siz*3)
			if FLIPCOIN() {
				b.r.X -= siz + spc
				if b.r.Y >= CNT.Y {
					if FLIPCOIN() {
						b.r.Y += RF32(siz/2, sizPREV+UHALF)
					} else {
						b.r.Y -= RF32(siz/2, sizPREV+UHALF)
					}
				} else {
					b.r.Y += RF32(siz/2, sizPREV+UHALF)
				}
			} else {
				if b.r.Y >= CNT.Y {
					b.r.Y += siz + spc
				} else {
					if FLIPCOIN() {
						b.r.Y += siz + spc
					} else {
						b.r.Y -= siz + spc
					}
				}
				b.r.X -= RF32(siz/2, sizPREV+UHALF)
			}
			if b.r.X <= borderINNER.X+siz {
				num = 0
				break
			}
			b.r = R(b.r.X, b.r.Y, siz, siz)
			if cADDBLOK(b) {
				room = append(room, b)
			}
			siz = osiz

			num--
		}
	case 11: //CENTER BLOKS
		spc += UQTR
		b := BLOK{}
		w := RF32(siz*3, siz*5)
		ow := w
		b.r = RECFROMCNT(CNT, w, w)
		if cADDBLOK(b) {
			room = append(room, b)
		}
		ox, oy := b.r.X, b.r.Y
		num := RINT(4, 8)
		for num > 0 {
			x, y := ox, oy
			w = RF32(siz-UHALF, siz*3)
			choose := RINT(1, 5)
			switch choose {
			case 1:
				y -= w + spc
				x -= w + spc
				x += RF32(0, ow)
			case 2:
				y -= w + spc
				x += ow + spc
				y += RF32(0, ow)
			case 3:
				y += ow + spc
				x -= w + spc
				x += RF32(0, ow)
			case 4:
				y -= w + spc
				x -= w + spc
				y += RF32(0, ow)
			}
			b.r = R(x, y, w, w)
			if cADDBLOK(b) {
				room = append(room, b)
			}
			num--
		}
	case 10: //RANDOM CORNER L
		siz -= spc
		num := RINT(5, 9)
		b := BLOK{}
		b.r = R(borderINNER.X+siz*5, borderINNER.Y+siz, siz, siz)
		ox := b.r.X
		for num > 0 {
			b.r.Y += RF32(siz*3, siz*6)
			if b.r.Y >= borderINNER.Y+borderINNER.Height-siz*3 {
				b.r.Y = borderINNER.Y + siz
				b.r.X += siz * 8
				ox = b.r.X
			}
			if cADDBLOKBORDER(b, U2) {
				room = append(room, b)
			}
			b.r.X -= siz
			if cADDBLOK(b) {
				room = append(room, b)
			}
			b.r.Y += siz
			if cADDBLOK(b) {
				room = append(room, b)
			}
			if b.r.Y >= borderINNER.Y+borderINNER.Height-siz*3 {
				b.r.Y = borderINNER.Y + siz
				b.r.X += siz * 4
				if FLIPCOIN() {
					b.r.X += siz
				}
				ox = b.r.X
			}
			b.r.X = ox
			if b.r.X >= borderINNER.X+borderINNER.Width-siz*4 {
				num = 0
			}

			num--
		}
	case 9: //CENTER CONCENTRIC SQUARES
		siz -= spc
		b := BLOK{}
		b.r = RECFROMCNT(CNT, siz, siz)
		b.r.Y -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X += siz * 6
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.Y += siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.Y += siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.Y -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.Y -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.Y -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X += siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X += siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X += siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X += siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.Y += siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.Y += siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.Y += siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.Y += siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.X -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}
		b.r.Y -= siz * 3
		if cADDBLOK(b) {
			room = append(room, b)
		}

	case 8: //DIAG RIGHT UP
		siz += UHALF
		x := borderINNER.X + siz
		y := borderINNER.Y + borderINNER.Height - siz*2
		oy := y
		ox := x
		b := BLOK{}
		for x < borderINNER.X+borderINNER.Width-(siz*5) {
			for y > borderINNER.Y+siz {
				if ROLL6() != 6 {
					b.r = R(x, y, siz, siz)
					if cADDBLOK(b) {
						room = append(room, b)
					}
				}
				x += siz
				y -= siz
				if x+siz >= borderINNER.X+borderINNER.Width {
					break
				}
			}
			y = oy
			x = ox
			x += siz * 5
			ox = x
		}
	case 7: //DIAG RIGHT DOWN
		siz += UHALF
		x := borderINNER.X + siz
		y := borderINNER.Y + siz
		oy := y
		ox := x
		b := BLOK{}
		for x < borderINNER.X+borderINNER.Width-(siz*5) {
			for y < borderINNER.Y+borderINNER.Height-siz {
				if ROLL6() != 6 {
					b.r = R(x, y, siz, siz)
					if cADDBLOK(b) {
						room = append(room, b)
					}
				}
				x += siz
				y += siz
				if x+siz >= borderINNER.X+borderINNER.Width {
					break
				}
			}
			y = oy
			x = ox
			x += siz * 5
			ox = x
		}
	case 6: //HORIZ LINES
		siz += UHALF
		x := borderINNER.X
		y := borderINNER.Y + siz
		b := BLOK{}
		for y+siz < borderINNER.Y+borderINNER.Height {
			x += RF32(0, borderINNER.Width/3)
			num := RINT(4, 12)
			for num > 0 {
				b.r = R(x, y, siz, siz)
				if cADDBLOK(b) {
					room = append(room, b)
				}
				x += siz
				num--
			}
			y += siz * 3
			if FLIPCOIN() {
				y += siz
			}
			x = borderINNER.X
		}
	case 5: //VERT LINES
		siz += U1
		x := borderINNER.X + siz*2
		y := borderINNER.Y + siz
		b := BLOK{}
		for x+siz < borderINNER.X+borderINNER.Width {
			y += RF32(0, borderINNER.Height/3)
			num := RINT(3, 6)
			for num > 0 {
				b.r = R(x, y, siz, siz)
				if cADDBLOK(b) {
					room = append(room, b)
				}
				y += siz
				num--
			}
			x += siz * 3
			if FLIPCOIN() {
				x += siz

			}
			y = borderINNER.Y
		}
	case 4: //ARROWS UPDOWN
		x := borderINNER.X + siz*2
		y := borderINNER.Y + siz*2
		b := BLOK{}
		for x+siz*3 < borderINNER.X+borderINNER.Width {
			UD := FLIPCOIN()
			if UD {
				y += siz + spc
			}
			for y < borderINNER.Y+borderINNER.Height {
				if UD {
					b.r = R(x+siz, y-siz, siz, siz)
				} else {
					b.r = R(x+siz, y, siz, siz)
				}
				if cADDBLOK(b) {
					room = append(room, b)
				}
				b.r.X -= siz + spc
				if UD {
					b.r.Y += siz + spc
				} else {
					b.r.Y -= siz + spc
				}
				if cADDBLOK(b) {
					room = append(room, b)
				}
				b.r.X += (siz * 2) + (spc * 2)
				if cADDBLOK(b) {
					room = append(room, b)
				}
				y += siz * 4
			}
			x += siz * 5
			y = borderINNER.Y + siz*2
		}
	case 3: //CROSS RAND
		x := borderINNER.X + siz*4
		b := BLOK{}
		for x+siz*3 < borderINNER.X+borderINNER.Width {
			if FLIPCOIN() {
				siz += U1
			}
			y := RF32(borderINNER.Y, borderINNER.Y+borderINNER.Height-siz*3)
			x2 := x - siz
			y2 := y + siz
			b.r = R(x2+siz*2, y2, siz, siz)
			if cADDBLOK(b) {
				room = append(room, b)
			}
			b.r = R(x, y, siz, siz)
			if cADDBLOK(b) {
				room = append(room, b)
			}
			b.r.Y += siz
			if cADDBLOK(b) {
				room = append(room, b)
			}
			b.r.Y += siz
			if cADDBLOK(b) {
				room = append(room, b)
			}
			b.r = R(x2, y2, siz, siz)
			if cADDBLOK(b) {
				room = append(room, b)
			}
			x += siz * 5
			if FLIPCOIN() {
				x += siz
			}
			siz = osiz
		}
	case 2: //CROSS ROWS
		x := borderINNER.X + siz*2
		c := 0
		for x < borderINNER.X+borderINNER.Width {
			y := borderINNER.Y
			b := BLOK{}
			for y+siz <= borderINNER.Y+borderINNER.Height {
				if c != 1 {
					x2 := x - siz
					y2 := y + siz
					b.r = R(x2+siz*2, y2, siz, siz)
					if !z.CheckCollisionRecs(b.r, startREC) {
						room = append(room, b)
					}
					b.r = R(x, y, siz, siz)
					if !z.CheckCollisionRecs(b.r, startREC) {
						room = append(room, b)
					}
					b.r.Y += siz
					if !z.CheckCollisionRecs(b.r, startREC) {
						room = append(room, b)
					}
					b.r.Y += siz
					if !z.CheckCollisionRecs(b.r, startREC) {
						room = append(room, b)
					}
					b.r = R(x2, y2, siz, siz)
					if !z.CheckCollisionRecs(b.r, startREC) {
						room = append(room, b)
					}
				}
				y += siz * 5

				c++
			}
			x += siz * 5
		}

	case 1: //COLUMN ROWS RND SIZ
		x := borderINNER.X + siz
		for x < (borderINNER.X+borderINNER.Width)-siz {
			y := borderINNER.Y
			b := BLOK{}
			for y+siz <= borderINNER.Y+borderINNER.Height {
				lrgr := FLIPCOIN()
				if lrgr {
					siz += U1
				}
				b.r = R(x, y, siz, siz)
				siz = osiz
				if !z.CheckCollisionRecs(b.r, startREC) {
					room = append(room, b)
				}
				if lrgr {
					y += siz * 4
				} else {
					y += siz * 3
				}
			}
			if FLIPCOIN() {
				x += siz * 3
			} else {
				x += siz * 4
			}
		}
	case 0: //COLUMN ROWS
		x := borderINNER.X + siz
		for x < borderINNER.X+borderINNER.Width {
			y := borderINNER.Y
			b := BLOK{}
			for y+siz <= borderINNER.Y+borderINNER.Height {
				b.r = R(x, y, siz, siz)
				if !z.CheckCollisionRecs(b.r, startREC) {
					room = append(room, b)
				}
				y += siz * 3
			}
			x += siz * 3
		}
	}

	mLEVEXTRAS()
}

//MARK: FIND
func fRECPOSITION(r z.Rectangle) z.Rectangle {
	r = RECFROMCNT(RECCNT(r), r.Width, r.Height)

	for i := range len(room) {
		if cRECREC(room[i].r, r) {
			r.X = room[i].r.X + room[i].r.Width
			r.Y = room[i].r.Y
			break
		}
	}

	return r
}
func fRECUDBORDEROUTER(r z.Rectangle) z.Rectangle {
	r.X = RF32(borderR.X+U1, borderR.X+borderR.Width-(U1+r.Width))
	if FLIPCOIN() {
		r.Y = borderR.Y
	} else {
		r.Y = borderR.Y + borderR.Height - r.Height
	}
	return r
}

//MARK: CHECK
func cADDREC(r z.Rectangle) bool {
	canadd := true
	if cRECREC(r, startREC) {
		canadd = false
	}
	if canadd {
		if len(room) > 0 {
			for i := range len(room) {
				if cRECREC(room[i].r, r) {
					canadd = false
					break
				}
			}
		}
	}

	if canadd {
		v1 := V2(r.X, r.Y)
		v2 := v1
		v2.X += r.Width
		v3 := v2
		v3.Y += r.Height
		v4 := v3
		v4.X -= r.Width
		if !z.CheckCollisionPointRec(v1, borderINNER) || !z.CheckCollisionPointRec(v2, borderINNER) || !z.CheckCollisionPointRec(v3, borderINNER) || !z.CheckCollisionPointRec(v4, borderINNER) {
			canadd = false
		}
	}

	return canadd
}
func cRECMOVE(r z.Rectangle) bool {
	canmove := true
	checkV2 := RECPOINTS(r)
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
func cADDBLOK(b BLOK) bool {
	canadd := true
	if cRECREC(b.r, startREC) {
		canadd = false
	}
	if canadd {
		if len(room) > 0 {
			for i := range len(room) {
				if cRECREC(room[i].r, b.r) {
					canadd = false
					break
				}
			}
		}
	}

	if canadd {
		v1 := V2(b.r.X, b.r.Y)
		v2 := v1
		v2.X += b.r.Width
		v3 := v2
		v3.Y += b.r.Height
		v4 := v3
		v4.X -= b.r.Width
		if !z.CheckCollisionPointRec(v1, borderINNER) || !z.CheckCollisionPointRec(v2, borderINNER) || !z.CheckCollisionPointRec(v3, borderINNER) || !z.CheckCollisionPointRec(v4, borderINNER) {
			canadd = false
		}
	}

	return canadd
}
func cADDBLOKBORDER(b BLOK, borderSIZ float32) bool {
	canadd := true
	if cRECREC(b.r, startREC) {
		canadd = false
	}
	if canadd {
		if len(room) > 0 {
			for i := range len(room) {
				if cRECREC(room[i].r, b.r) {
					canadd = false
					break
				}
			}
		}
	}
	if canadd {
		if len(room) > 0 {
			for i := range len(room) {
				cR1, cR2, cR3, cR4 := b.r, b.r, b.r, b.r
				cR1.Y -= borderSIZ
				cR2.X += borderSIZ
				cR3.Y += borderSIZ
				cR4.X -= borderSIZ
				if cRECREC(room[i].r, cR1) || cRECREC(room[i].r, cR2) || cRECREC(room[i].r, cR3) || cRECREC(room[i].r, cR4) {
					canadd = false
					break
				}
			}
		}
	}
	if canadd {
		v1 := V2(b.r.X, b.r.Y)
		v2 := v1
		v2.X += b.r.Width
		v3 := v2
		v3.Y += b.r.Height
		v4 := v3
		v4.X -= b.r.Width
		if !z.CheckCollisionPointRec(v1, borderINNER) || !z.CheckCollisionPointRec(v2, borderINNER) || !z.CheckCollisionPointRec(v3, borderINNER) || !z.CheckCollisionPointRec(v4, borderINNER) {
			canadd = false
		}
	}
	return canadd
}

//MARK: UTILS
func OBJSPD(o OBJ) OBJ {
	for {
		o.velx = RF32(-SIZMULTIDIV(o.maxSPD, 20, 16), SIZMULTIDIV(o.maxSPD, 20, 16))
		if ABS(o.velx) >= o.maxSPD/3 {
			break
		}
	}
	for {
		o.vely = RF32(-SIZMULTIDIV(o.maxSPD, 20, 16), SIZMULTIDIV(o.maxSPD, 20, 16))
		if ABS(o.vely) >= o.maxSPD/3 {
			break
		}
	}
	return o
}
