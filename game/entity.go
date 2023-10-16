package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type IEntity interface {
	Type() string
	Update()
	Draw(screen *ebiten.Image)
	X() float64
	Y() float64
}

type Entity struct {
	op     *ebiten.DrawImageOptions
	img    *ebiten.Image
	state  *State
	x, y   float64
	vx, vy float64
	speed  float64
	scale  float64
}

func (e *Entity) Draw(screen *ebiten.Image) {
	e.op.GeoM.Reset()
	e.op.GeoM.Scale(e.scale, e.scale)
	e.op.GeoM.Translate(e.x, e.y)
	screen.DrawImage(e.img, e.op)
}
