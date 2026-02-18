package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	z.InitWindow(SCRW, SCRH, "TITLE HERE")

	INIT()

	for !z.WindowShouldClose() {
		FRAMES++
		z.BeginDrawing()

		z.ClearBackground(z.Black)
		if z.IsKeyPressed(z.KeyF2) {
			z.ClearBackground(z.Black)
		}
		//DRAW CAM SHADER
		if SHADER1ON || SHADER2ON {
			z.BeginTextureMode(RNDRTEX)
			z.ClearBackground(z.Black)
		}
		z.BeginMode2D(CAM)
		dCAM()
		if !STARTED {
			STARTED = true
		}
		z.EndMode2D()
		if SHADER1ON || SHADER2ON {
			z.EndTextureMode()
		}
		//END DRAW CAM SHADER

		if SHADER1ON && !SHADER2ON {
			z.BeginShaderMode(SHADER1)
			z.DrawTextureRec(RNDRTEX.Texture, z.NewRectangle(0, 0, float32(RNDRTEX.Texture.Width), float32(-RNDRTEX.Texture.Height)), z.NewVector2(0, 0), z.White)
			z.EndShaderMode()
		} else if !SHADER1ON && SHADER2ON {
			z.BeginShaderMode(SHADER2)
			z.DrawTextureRec(RNDRTEX.Texture, z.NewRectangle(0, 0, float32(RNDRTEX.Texture.Width), float32(-RNDRTEX.Texture.Height)), z.NewVector2(0, 0), z.White)
			z.EndShaderMode()
		} else if SHADER1ON && SHADER2ON {
			z.BeginTextureMode(RNDRTEX2)
			z.BeginShaderMode(SHADER2)
			z.DrawTextureRec(RNDRTEX.Texture, z.NewRectangle(0, 0, float32(RNDRTEX.Texture.Width), float32(-RNDRTEX.Texture.Height)), z.NewVector2(0, 0), z.White)
			z.EndShaderMode()
			z.EndTextureMode()

			z.BeginShaderMode(SHADER1)
			z.DrawTextureRec(RNDRTEX2.Texture, z.NewRectangle(0, 0, float32(RNDRTEX2.Texture.Width), float32(-RNDRTEX2.Texture.Height)), z.NewVector2(0, 0), z.White)
			z.EndShaderMode()
		}

		//DRAW NOCAM
		dNOCAM()
		//ENDDRAW NOCAM
		z.EndDrawing()
		UP()
	}

	EXIT()

	z.CloseWindow()
}
