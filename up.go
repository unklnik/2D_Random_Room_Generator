package main

import (
	z "github.com/gen2brain/raylib-go/raylib"
)

var (
	plexitNUM   = BLANKNUM
	nextLEV     bool
	nextROOMNUM int
	roomchangeT int32
	delta       float32
)

func UP() {
	delta = z.GetFrameTime()
	delta *= 60
	INP()
	uPL()
	uCAM()
	if STARTED {

		if len(roomOBJ) > 0 {
			cOBJSTART()
			uOBJ()
		}
	}

	TIMERS()

	if roomchangeT == 0 {
		if nextLEV {

		}
	}
}

// MARK: CAM
func uCAM() {

}

// MARK: TIMERS
func TIMERS() {

	if roomchangeT > 0 {
		roomchangeT--
	}

	if fadeliteONOFF {
		if fadeALITE >= fadeMINLITE {
			fadeALITE -= fadeSPD
		} else {
			fadeliteONOFF = false
		}
	} else {
		if fadeALITE <= fadeMAXLITE {
			fadeALITE += fadeSPD
		} else {
			fadeliteONOFF = true
		}
	}
	if fadeONOFF {
		if fadeA >= fadeMIN {
			fadeA -= fadeSPD
		} else {
			fadeONOFF = false
		}
	} else {
		if fadeA <= fadeMAX {
			fadeA += fadeSPD
		} else {
			fadeONOFF = true
		}
	}

}
