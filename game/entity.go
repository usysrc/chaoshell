package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Type() string
	Update()
	Draw(screen *ebiten.Image)
}
