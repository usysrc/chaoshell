package component

import "github.com/hajimehoshi/ebiten/v2"

type Position struct {
	X, Y float64
}

type Velocity struct {
	X, Y float64
}

type Render struct {
	Image *ebiten.Image
	Z     float64
	Scale float64
}
