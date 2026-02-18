package main

import (
	"fmt"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	TEX     []z.Texture2D
	BLANKIM IM

	TILES, ETCIM, PLIM, BGIM, INNERIM []IM

	EXITARROW, EXITARROW1, EXITARROW2, EXITARROW3, EXITARROW4 ANIM

	floorT, wallT, innerT IM
)

type IM struct {
	tex int
	r   z.Rectangle
	c   z.Color
	ro  float32
}
type ANIM struct {
	ims                       []IM
	frameT                    int32
	totalFrames, currentFrame int
	complete                  bool
}

// MARK: MAKE
func mIMGS() {
	mETCIM("res/etc")
	BGIM = mIMSHEETpath("res/bg/")
	TILES = mIMSHEETpath("res/tiles/")
	INNERIM = mIMSHEETpath("res/inner/")
	exitanim := MIMSHEETIM(ETCIM[4], 10, 1, 0, 0, 48, 48, 0, 0)
	EXITARROW = mANIMIMSHEET(exitanim, 15, 10)
	EXITARROW1, EXITARROW2, EXITARROW3, EXITARROW4 = EXITARROW, EXITARROW, EXITARROW, EXITARROW

	floorT = BGIM[RINT(0, len(BGIM))]
	wallT = TILES[RINT(0, len(TILES))]
	innerT = INNERIM[RINT(0, len(INNERIM))]

}

// MARK: DRAW ANIM
func dANIM(anm ANIM, r z.Rectangle) ANIM {
	dIM(anm.ims[anm.currentFrame], r, z.White)
	if FRAMES%anm.frameT == 0 {
		anm.currentFrame++
	}
	if anm.currentFrame == anm.totalFrames {
		anm.currentFrame = 0
	}
	return anm
}
func dANIMANGL(anm ANIM, r z.Rectangle, angl float32) ANIM {
	dIMRO(anm.ims[anm.currentFrame], r, z.White, angl)
	if FRAMES%anm.frameT == 0 {
		anm.currentFrame++
	}
	if anm.currentFrame == anm.totalFrames {
		anm.currentFrame = 0
	}
	return anm
}
func dANIMANGLSHADOW(anm ANIM, r z.Rectangle, angl, offset float32, cShadow z.Color) ANIM {
	dIMROSHADOW(anm.ims[anm.currentFrame], r, z.White, cShadow, angl, offset)
	if FRAMES%anm.frameT == 0 {
		anm.currentFrame++
	}
	if anm.currentFrame == anm.totalFrames {
		anm.currentFrame = 0
	}
	return anm
}
func dANIMONCE(anm ANIM, r z.Rectangle) ANIM {
	dIM(anm.ims[anm.currentFrame], r, z.White)
	if FRAMES%anm.frameT == 0 {
		anm.currentFrame++
	}
	if anm.currentFrame == anm.totalFrames {
		anm.complete = true
	}
	return anm
}
func dANIMFLIP(anm ANIM, r z.Rectangle) ANIM {
	dIMFLIP(anm.ims[anm.currentFrame], r)
	if FRAMES%anm.frameT == 0 {
		anm.currentFrame++
	}
	if anm.currentFrame == anm.totalFrames {
		anm.currentFrame = 0
	}
	return anm
}
func dANIMCOLOR(anm ANIM, r z.Rectangle, c z.Color) ANIM {
	dIM(anm.ims[anm.currentFrame], r, c)
	if FRAMES%anm.frameT == 0 {
		anm.currentFrame++
	}
	if anm.currentFrame == anm.totalFrames {
		anm.currentFrame = 0
	}
	return anm
}
func dANIMCOLORFLIP(anm ANIM, r z.Rectangle, c z.Color) ANIM {
	dIMFLIPCOLOR(anm.ims[anm.currentFrame], r, c)
	if FRAMES%anm.frameT == 0 {
		anm.currentFrame++
	}
	if anm.currentFrame == anm.totalFrames {
		anm.currentFrame = 0
	}
	return anm
}

// MARK: MAKE ANIM
func mANIMIMSHEET(ims []IM, fps int32, numFrames int) ANIM {
	anm := ANIM{}
	anm.frameT = FPS / fps
	anm.totalFrames = numFrames
	for i := range numFrames {
		anm.ims = append(anm.ims, ims[i])
	}
	return anm
}

// MARK: DRAW IM
func dIMROSHADOW(im IM, r z.Rectangle, c, cShadow z.Color, angl, offset float32) {
	r2 := r
	r2.X -= offset
	r2.Y += offset
	z.DrawTexturePro(TEX[im.tex], im.r, DR(r2), ORIG(r2), angl, cShadow)
	z.DrawTexturePro(TEX[im.tex], im.r, DR(r), ORIG(r), angl, c)
}
func dIMRO(im IM, r z.Rectangle, c z.Color, angl float32) {
	z.DrawTexturePro(TEX[im.tex], im.r, DR(r), ORIG(r), angl, c)
}
func dIMROTATES(im IM, r z.Rectangle, c z.Color, roSPEED float32) IM {
	z.DrawTexturePro(TEX[im.tex], im.r, DR(r), ORIG(r), im.ro, c)
	im.ro += roSPEED
	return im
}
func dIMROTATESSHADOW(im IM, r z.Rectangle, c, cShadow z.Color, offset, roSPEED float32) IM {
	r2 := r
	r2.Y += offset
	r2.X -= offset
	z.DrawTexturePro(TEX[im.tex], im.r, DR(r2), ORIG(r2), im.ro, cShadow)
	z.DrawTexturePro(TEX[im.tex], im.r, DR(r), ORIG(r), im.ro, c)
	im.ro += roSPEED
	return im
}
func dIM(im IM, r z.Rectangle, c z.Color) {
	z.DrawTexturePro(TEX[im.tex], im.r, r, z.Vector2Zero(), 0, c)
}
func dIMSHADOW(im IM, r z.Rectangle, c, cShadow z.Color, offset float32) {
	r2 := r
	r2.X -= offset
	r2.Y += offset
	z.DrawTexturePro(TEX[im.tex], im.r, r2, z.Vector2Zero(), 0, cShadow)
	z.DrawTexturePro(TEX[im.tex], im.r, r, z.Vector2Zero(), 0, c)
}
func dIMFLIPCOLOR(im IM, r z.Rectangle, c z.Color) {
	im.r.Width = -im.r.Width
	z.DrawTexturePro(TEX[im.tex], im.r, r, z.Vector2Zero(), 0, c)
}
func dIMFLIP(im IM, r z.Rectangle) {
	im.r.Width = -im.r.Width
	z.DrawTexturePro(TEX[im.tex], im.r, r, z.Vector2Zero(), 0, z.White)
}
func dIMSHEET(ims []IM, x, y, space, zoom float32) {
	ox := x
	for i := range ims {
		z.DrawTexturePro(TEX[ims[i].tex], ims[i].r, R(x, y, ims[i].r.Width*zoom, ims[i].r.Height*zoom), z.Vector2Zero(), 0, z.White)
		dTXT1XY(fmt.Sprint(i), x, y+ims[i].r.Height*zoom+2)
		x += space + ims[i].r.Width*zoom
		if x+ims[i].r.Width*zoom >= float32(SCRW) {
			x = ox
			y += space + ims[i].r.Height*zoom + F1H + 4
		}

	}
}
func dIMSHEETSCROLL(ims []IM, x, y, space, zoom, yChange float32) float32 {
	ox := x
	y += yChange
	for i := range ims {
		z.DrawTexturePro(TEX[ims[i].tex], ims[i].r, R(x, y, ims[i].r.Width*zoom, ims[i].r.Height*zoom), z.Vector2Zero(), 0, z.White)
		dTXT1XY(fmt.Sprint(i), x, y+ims[i].r.Height*zoom+2)
		x += space + ims[i].r.Width*zoom
		if x+ims[i].r.Width*zoom >= float32(SCRW) {
			x = ox
			y += space + ims[i].r.Height*zoom + F1H + 4
		}
		yChange = dSCROLLVERT(yChange, U1, z.Orange)

	}
	return yChange
}

// MARK: MAKE
func mIMSHEETpath(path string) []IM {
	var ims []IM
	images := cPNG(path)
	println(len(images))
	for i := range len(images) {
		ims = append(ims, mIMfile(images[i]))
	}
	return ims
}
func mETCIM(path string) {
	images := cPNG(path)
	println(len(images))
	for i := range len(images) {
		ETCIM = append(ETCIM, mIMfile(images[i]))
	}
}
func mIMfile(path string) IM {
	im := IM{}
	image := z.LoadImage(path)
	TEX = append(TEX, z.LoadTextureFromImage(image))
	w := float32(image.Width)
	h := float32(image.Height)
	z.UnloadImage(image)
	im.r = R(0, 0, w, h)
	im.tex = len(TEX) - 1
	im.c = gameColor
	return im
}
func mIM(path string, x, y, w, h float32) IM {
	im := IM{}
	image := z.LoadImage(path)
	TEX = append(TEX, z.LoadTextureFromImage(image))
	z.UnloadImage(image)
	im.r = R(x, y, w, h)
	im.tex = len(TEX) - 1
	return im
}
func MIMSHEETIM(im IM, numW, numH int, x, y, w, h, offsetX, offsetY float32) []IM {
	var ims []IM
	a := numW * numH
	c := 0
	ox := x
	for a > 0 {
		im2 := IM{}
		im2.r = R(x, y, w, h)
		im2.tex = im.tex
		ims = append(ims, im2)
		x += w + offsetX
		c++
		a--
		if c == numW {
			c = 0
			x = ox
			y += h + offsetY
		}
	}
	return ims
}
func mIMSHEETXY(path string, numW, numH int, x, y, w, h, offsetX, offsetY float32) []IM {
	var ims []IM
	a := numW * numH
	c := 0
	ox := x
	image := z.LoadImage(path)
	TEX = append(TEX, z.LoadTextureFromImage(image))
	z.UnloadImage(image)
	for a > 0 {
		im := IM{}
		im.r = R(x, y, w, h)
		im.tex = len(TEX) - 1
		ims = append(ims, im)
		x += w + offsetX
		c++
		a--
		if c == numW {
			c = 0
			x = ox
			y += h + offsetY
		}
	}

	return ims
}
