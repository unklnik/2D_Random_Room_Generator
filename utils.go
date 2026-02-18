package main

import (
	"math"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"

	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	BLANKNUM = 7777777777777777777
)

// MARK: TRI
func cPOINTTRI(v2 z.Vector2, t TRI) bool {
	return z.CheckCollisionPointTriangle(v2, t.v2[0], t.v2[1], t.v2[2])
}

// MARK: RECS
func RECPOINTS(r z.Rectangle) []z.Vector2 {
	v1 := V2(r.X, r.Y)
	v2 := v1
	v2.X += r.Width
	v3 := v2
	v3.Y += r.Height
	v4 := v3
	v4.X -= r.Width
	return []z.Vector2{v1, v2, v3, v4}
}
func RESIZERECWIDTH(r z.Rectangle, newWidth float32) z.Rectangle {
	aspect := r.Height / r.Width
	newHeight := newWidth * aspect
	return z.Rectangle{X: r.X, Y: r.Y, Width: newWidth, Height: newHeight}
}
func DR(r z.Rectangle) z.Rectangle {
	return z.NewRectangle(r.X+r.Width/2, r.Y+r.Height/2, r.Width, r.Height)
}
func RECSLICERANDARRANGE(r []LEVREC) []LEVREC {
	var r2 []LEVREC
	r2 = append(r2, r[0])
	for i := 1; i < len(r); i++ {
		for {
			choose := 0
			if len(r2) > 2 {
				choose = RINT(0, len(r2))
			}
			r[i].r.X = r[choose].r.X
			r[i].r.Y = r[choose].r.Y

			chooseSIDE := RINT(1, 5)
			switch chooseSIDE {
			case 1:
				r[i].r.Y -= r[i].r.Height
			case 2:
				r[i].r.X += r[i].r.Width
			case 3:
				r[i].r.Y += r[i].r.Height
			case 4:
				r[i].r.X -= r[i].r.Width
			}
			collides := false
			for j := range r2 {
				if z.CheckCollisionRecs(r[i].r, r2[j].r) {
					collides = true
					break
				}
			}
			if !collides {
				r[i].num = i
				r2 = append(r2, r[i])
				break
			}
		}
	}

	FINDRECADJACENTEXITS(r2)
	return r2
}

func FINDRECADJACENTEXITS(level []LEVREC) {
	const eps = 1e-6
	for i := range level {
		sides := []int{}
		r := level[i].r
		for j := range level {
			if i == j {
				continue
			}
			other := level[j].r
			if math.Abs(float64(r.Y-(other.Y+other.Height))) < eps {
				if r.X < other.X+other.Width && r.X+r.Width > other.X {
					sides = append(sides, 1)
				}
			}
			if math.Abs(float64((r.X+r.Width)-other.X)) < eps {
				if r.Y < other.Y+other.Height && r.Y+r.Height > other.Y {
					sides = append(sides, 2)
				}
			}
			if math.Abs(float64((r.Y+r.Height)-other.Y)) < eps {
				if r.X < other.X+other.Width && r.X+r.Width > other.X {
					sides = append(sides, 3)
				}
			}
			if math.Abs(float64(r.X-(other.X+other.Width))) < eps {
				if r.Y < other.Y+other.Height && r.Y+r.Height > other.Y {
					sides = append(sides, 4)
				}
			}
		}
		level[i].exits = uniqueInts(sides)
	}
}

func uniqueInts(in []int) []int {
	seen := make(map[int]bool)
	out := []int{}
	for _, v := range in {
		if !seen[v] {
			seen[v] = true
			out = append(out, v)
		}
	}
	return out
}

func RECCNT(r z.Rectangle) z.Vector2 {
	return z.NewVector2(r.X+r.Width/2, r.Y+r.Height/2)
}
func RECFROMCNT(cnt z.Vector2, w, h float32) z.Rectangle {
	return z.NewRectangle(cnt.X-w/2, cnt.Y-h/2, w, h)
}
func RECSMLR(r z.Rectangle, offset float32) z.Rectangle {
	return R(r.X+offset, r.Y+offset, r.Width-offset*2, r.Height-offset*2)
}
func RECLRGR(r z.Rectangle, offset float32) z.Rectangle {
	return R(r.X-offset, r.Y-offset, r.Width+offset*2, r.Height+offset*2)
}
func RECSHIFT2DIR(r z.Rectangle, v2 z.Vector2) z.Rectangle {
	r.X += v2.X
	r.Y += v2.Y
	return r
}
func RECSHIFTURDL(r z.Rectangle, offset float32, urdl1234 int) z.Rectangle {
	switch urdl1234 {
	case 1: //UP
		r.Y -= offset
	case 2: //RIGHT
		r.X += offset
	case 3: //DOWN
		r.Y += offset
	case 4: //LEFT
		r.X -= offset
	}
	return r
}

// MARK:COLLIS
func cPOINTREC(v2 z.Vector2, r z.Rectangle) bool {
	return z.CheckCollisionPointRec(v2, r)
}
func cPOINTBLOKSCLICE(v2 z.Vector2, b []BLOK) int {
	num := 0
	for i := range b {
		if z.CheckCollisionPointRec(v2, b[i].r) {
			num = i
			break
		}
	}
	return num
}

func cRECBLOKSLICE(r z.Rectangle, b []BLOK) bool {
	collides := false
	for i := range b {
		if z.CheckCollisionRecs(r, b[i].r) {
			collides = true
			break
		}
	}
	return collides
}
func cRECRECSLICE(r z.Rectangle, rs []z.Rectangle) bool {
	collides := false
	for i := range rs {
		if z.CheckCollisionRecs(r, rs[i]) {
			collides = true
			break
		}
	}
	return collides
}
func cRECREC(r, r2 z.Rectangle) bool {
	return z.CheckCollisionRecs(r, r2)
}

func cBLOKRECINCIRC(b BLOK, cnt z.Vector2, rad float32) bool {
	collides := false
	if z.CheckCollisionCircleRec(cnt, rad, b.r) {
		collides = true
	}
	return collides
}

// MARK:VECTOR2
func ORIG(r z.Rectangle) z.Vector2 {
	return z.NewVector2(r.Width/2, r.Height/2)
}
func MOVEV2SLICE(v2 []z.Vector2, dirx, diry float32) []z.Vector2 {
	for i := range len(v2) {
		v2[i].X += dirx
		v2[i].Y += diry
	}
	return v2
}
func MOVEV2(v2 z.Vector2, dirx, diry float32) z.Vector2 {
	v2.X += dirx
	v2.Y += diry
	return v2
}
func V2(x, y float32) z.Vector2 {
	return z.NewVector2(x, y)
}
func NORMALIZE(v z.Vector2) z.Vector2 {
	length := float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
	if length == 0 {
		return z.NewVector2(0, 0)
	}
	return z.NewVector2(v.X/length, v.Y/length)
}

// MARK: FILES
func cPNG(path string) []string {
	var pngFiles []string

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info != nil && !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".png") {
			pngFiles = append(pngFiles, path)
		}
		return nil
	})
	return pngFiles
}

// MARK: SLICES
func SLICECONTAINS[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}
func REMSLICESTRUCT[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s // unchanged if out of range
	}
	return append(s[:i], s[i+1:]...)
}
func TRIMSLICESTRUC[T any](s []T, n int) []T {
	if n < 0 {
		return []T{}
	}
	if n > len(s) {
		return s
	}
	return s[:n]
}

// MARK: MATH
func ABS(num float32) float32 {
	return float32(math.Abs(float64(num)))
}
func ABSDIFF(a, b float32) float32 { return float32(math.Abs(float64(a - b))) }
func SIZMULTIDIV(siz, divNum, multiNum float32) float32 {
	return (siz / divNum) * multiNum
}
func U1DIV(div float32) float32 {
	return U1 / div
}
func U1MULTIDIV(div, multi float32) float32 {
	return (U1 / div) * multi
}

// MARK: RANDOM
func ROLL6() int {
	return rand.IntN(6) + 1
}
func ROLL12() int {
	return rand.IntN(12) + 1
}
func FLIPCOIN() bool {
	return rand.IntN(2) == 0
}
func RINT(min, max int) int {
	return min + rand.IntN(max-min)
}
func RF32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
func RUINT8(min, max int) uint8 {
	return uint8(min) + uint8(rand.IntN(int(max-min+1)))
}
