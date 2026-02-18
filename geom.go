package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var ()

type BLOK struct {
	r     z.Rectangle
	c, c2 z.Color
	nm    string
	cnt   z.Vector2
	im    IM

	solid, off, activ bool

	lit int
}
type LEVREC struct {
	r                z.Rectangle
	num              int
	exits, nextROOMS []int
	visited          bool
}

type TRI struct {
	v2 []z.Vector2
}

// MARK: LEVEL ROOMS
func dINNER() {
	for i := range room {
		shadowOffset2 := SIZMULTIDIV(room[i].r.Width, 40, 3)
		if shadowOffset2 > shadowOffset {
			shadowOffset2 = shadowOffset
		}
		dIMSHADOW(innerT, room[i].r, gameColor, shadowColor, shadowOffset2)
		if DBG {
			dRECLINES(room[i].r, 4, BLUE())
		}
	}
}
func dFLOORWALLS() {

	//FLOORS
	dREC(borderR, BLACK())
	num := 10
	if floorT.r.Width < U2 {
		num = 5
	}
	w := borderR.Width / float32(num)
	x := borderR.X
	y := borderR.Y
	r := RESIZERECWIDTH(floorT.r, w)

	for y < borderR.Y+borderR.Height {
		if x >= borderR.X+borderR.Width {
			x = borderR.X
			y += r.Height
		}
		r.X = x
		r.Y = y
		dIM(floorT, r, bgColor)
		x += r.Width
	}

	//BORDER RECS TO COVER OVERLAP
	dRECXYF32(0, borderR.Y+borderR.Height, float32(SCRW), float32(SCRH)-borderR.Y+borderR.Height, BLACK())
	dRECXYF32(borderR.X+borderR.Width, 0, float32(SCRW)-borderR.Width, float32(SCRH), BLACK())

	//WALLS
	x = borderR.X - blokW
	y = borderR.Y - blokW
	for x < borderR.X+borderR.Width+blokW {
		dIM(wallT, R(x, y, blokW, blokW), gameColor)
		y2 := y + borderR.Height + blokW
		dIM(wallT, R(x, y2, blokW, blokW), gameColor)
		x += blokW
	}
	x = borderR.X - blokW
	y = borderR.Y
	for y < borderR.Y+borderR.Height {
		dIM(wallT, R(x, y, blokW, blokW), gameColor)
		x2 := x + borderR.Width + blokW
		dIM(wallT, R(x2, y, blokW, blokW), gameColor)
		y += blokW
	}
}

// MARK: LINES
func dLINEXY(x1, y1, x2, y2, lineW float32, c z.Color) {
	z.DrawLineEx(V2(x1, y1), V2(x2, y2), lineW, c)
}
func dLINEV2(v1, v2 z.Vector2, lineW float32, c z.Color) {
	z.DrawLineEx(v1, v2, lineW, c)
}

// MARK: TRIANGLES
func dTRI(t TRI, c z.Color) {
	z.DrawTriangle(t.v2[1], t.v2[0], t.v2[2], c)
}
func dTRILINES(t TRI, c z.Color) {
	z.DrawTriangleLines(t.v2[1], t.v2[0], t.v2[2], c)
}

func m2XTRI4POINTS(v2 []z.Vector2) []TRI {
	var tri []TRI
	t := TRI{}
	t.v2 = append(t.v2, v2[0], v2[1], v2[2])
	tri = append(tri, t)
	t = TRI{}
	t.v2 = append(t.v2, v2[0], v2[2], v2[3])
	tri = append(tri, t)
	return tri
}

// MARK:RECTANGLES
func R(x, y, w, h float32) z.Rectangle {
	return z.NewRectangle(x, y, w, h)
}
func dREC(r z.Rectangle, c z.Color) {
	z.DrawRectangleRec(r, c)

}
func dRECLINES(r z.Rectangle, lineW float32, c z.Color) {
	z.DrawRectangleLinesEx(r, lineW, c)
}
func dRECXY(x, y, w, h int32, c z.Color) {
	z.DrawRectangle(x, y, w, h, c)
}
func dRECXYF32(x, y, w, h float32, c z.Color) {
	x2 := int32(x)
	y2 := int32(y)
	w2 := int32(w)
	h2 := int32(h)
	z.DrawRectangle(x2, y2, w2, h2, c)
}
func dRECXYLINES(x, y, w, h int32, c z.Color) {
	z.DrawRectangleLines(x, y, w, h, c)
}
func dRECCNT(cnt z.Vector2, w, h int32, c z.Color) {
	z.DrawRectangle(int32(cnt.X)-w/2, int32(cnt.Y)-h/2, w, h, c)
}
func dRECCNTLINES(cnt z.Vector2, w, h int32, c z.Color) {
	z.DrawRectangleLines(int32(cnt.X)-w/2, int32(cnt.Y)-h/2, w, h, c)
}
