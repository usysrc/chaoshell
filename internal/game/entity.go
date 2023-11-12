package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var nextID int = 0

func getNextID() int {
	nextID++
	return nextID
}

type Entifier interface {
	Type() string
	ID() int
	Update()
	Draw(screen *ebiten.Image)
	X() float64
	Y() float64
}

type Entity struct {
	op        *ebiten.DrawImageOptions
	img       *ebiten.Image
	container *Container
	x, y      float64
	vx, vy    float64
	speed     float64
	scale     float64
	id        int
}

func (e *Entity) Update() {

}

func (e *Entity) Draw(screen *ebiten.Image) {
	e.op.GeoM.Reset()
	e.op.GeoM.Scale(e.scale, e.scale)
	e.op.GeoM.Translate(e.x, e.y)
	screen.DrawImage(e.img, e.op)
}

func (e *Entity) Type() string {
	return "entity"
}

func (e *Entity) ID() int {
	return e.id
}
