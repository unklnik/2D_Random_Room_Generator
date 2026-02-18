package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	SHADER1, SHADER2                z.Shader
	SHADER1ON, SHADER2ON, SCANLINES bool

	RNDRTEX, RNDRTEX2 z.RenderTexture2D

	DOOR0ANIM, DOOR1ANIM, DOOR2ANIM, DOOR3ANIM ANIM
)

// MARK: INIT
func INITSHADERS() {
	SHADER1 = z.LoadShader("", "res/blur.fs")
	SHADER2 = z.LoadShader("", "res/bloom.fs")

	RNDRTEX = z.LoadRenderTexture(SCRW*2, SCRH*2)
	RNDRTEX2 = z.LoadRenderTexture(SCRW*2, SCRH*2)
}

// MARK: DRAW
func dDOORS() {

}
func dSCAN(spc, lineW float32, c z.Color) {
	var x, y float32
	x2 := x + float32(SCRW)
	for y < float32(SCRH) {
		dLINEXY(x, y, x2, y, lineW, c)
		y += spc
	}
}
